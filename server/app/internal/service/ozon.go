package service

import (
	"context"
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/mapper"
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/common/pkg/utils"
	"net/http"
	"time"
)

type productInfo struct {
	Sku int64
	Barcode,
	BarcodeId,
	ItemId *string
	Items []*model.StockItem
}

type products struct {
	data map[int64]*productInfo
}

func newProducts() *products {
	return &products{
		data: make(map[int64]*productInfo),
	}
}

func (p *products) getSkus() *[]int64 {
	skus := make([]int64, len(p.data))
	index := 0
	for sku := range p.data {
		skus[index] = sku
		index++
	}
	return &skus
}

func (p *products) addProduct(sku int64, product *model.StockItem) bool {
	val, ok := p.data[sku]
	if !ok {
		val = &productInfo{
			Sku:       sku,
			Barcode:   product.Barcode,
			BarcodeId: product.BarcodeId,
			ItemId:    product.ItemId,
			Items:     []*model.StockItem{},
		}
		p.data[sku] = val
	}
	val.Items = append(val.Items, product)
	return true
}

func (p *products) get(sku int64) (*productInfo, bool) {
	val, ok := p.data[sku]
	return val, ok
}

type OzonService struct {
	clientId         string
	apiKey           string
	ownerCode        string
	productInfoCache map[int64]*string
	config           *configuration.Config
	client           *api.ClientWithResponses
	products         *products
}

func NewOzonService(ownerCode, clientId, apiKey string, config *configuration.Config) *OzonService {
	return NewOzonServiceWIthClient(ownerCode, clientId, apiKey, nil, config)
}

func NewOzonServiceWIthClient(ownerCode, clientId, apiKey string, client *api.ClientWithResponses, config *configuration.Config) *OzonService {
	return &OzonService{
		clientId:         clientId,
		apiKey:           apiKey,
		config:           config,
		ownerCode:        ownerCode,
		client:           client,
		productInfoCache: make(map[int64]*string),
		products:         newProducts(),
	}
}

func (c *OzonService) customProvider(ctx context.Context, req *http.Request) error {
	req.Header.Set("Client-Id", c.clientId)
	req.Header.Set("Api-Key", c.apiKey)
	return nil
}

func (c *OzonService) Parsing(context context.Context, transactionID int64) error {
	logrus.Info("Start parsing ozon ", c.ownerCode)
	client, err := c.getClient(c.config.Ozon.Host)
	if err != nil {
		return err
	}
	offset := 0
	whType := api.GetOzonSupplierStocksJSONBodyWarehouseTypeALL
	source := types.SourceTypeOzon
	dt := time.Now()
	//----------------- LOAD STOCK -----------------
	var data []*model.StockItem
	for {
		resp, err := client.GetOzonSupplierStocksWithResponse(context, api.GetOzonSupplierStocksJSONRequestBody{
			Limit:         &c.config.BatchSize,
			Offset:        &offset,
			WarehouseType: &whType,
		})
		if err != nil {
			return err
		}
		if resp.StatusCode() != 200 {
			return errors.New(fmt.Sprintf("Ozon resp %d: %s", resp.StatusCode(), resp.Status()))
		}
		items := resp.JSON200.Result.Rows
		newItem, err := c.createItems(items, dt, transactionID, source)
		if err != nil {
			return err
		}
		data = append(data, *newItem...)
		length := len(*items)
		if length < c.config.BatchSize {
			break
		}
		offset += length
	}
	logrus.Info("Take an information about prices")
	err = c.enrichProductPrices(client, context, source)
	storage.SaveStocksInBatches(&data, c.config.BatchSize)
	if err != nil {
		return err
	}
	//------------------ LOAD ORDERS ------------------
	logrus.Info("Load orders")
	if err := c.loadOrders(context, client, transactionID, source); err != nil {
		return err
	}
	logrus.Info("Finish to load orders")
	logrus.Info("Finish parsing ozon")
	return nil
}

func (c *OzonService) enrichProductPrices(client *api.ClientWithResponses, ctx context.Context, source string) error {
	skus := c.products.getSkus()

	decoder := dictionary.GetItemDecoder()
	err := CallbackBatch[int64](skus, c.config.BatchSize, func(batch *[]int64) error {
		req := api.GetOzonProductInfoJSONRequestBody{
			Sku: batch,
		}
		timeCtx, _ := context.WithTimeout(ctx, time.Duration(c.config.Timeout)*time.Second)
		resp, err := client.GetOzonProductInfoWithResponse(timeCtx, req)
		if err != nil {
			return err
		}
		if resp.StatusCode() != 200 || resp.JSON200 == nil {
			return errors.New(fmt.Sprintf("Ozon resp %d %s Body: %s", resp.StatusCode(), resp.Status(), string(resp.Body)))
		}
		for _, info := range *resp.JSON200.Result.Items {
			item, ok := c.products.get(*info.FboSku)
			if !ok {
				logrus.Errorf("Can't find product with sku %d", *info.FboSku)
				continue
			}
			item.Barcode = info.Barcode
			var barcodeId, itemId, message *string
			if info.Barcode != nil {
				decode, err := decoder.Decode(c.ownerCode, source, *info.Barcode)
				if err != nil {
					s := err.Error()
					message = &s
				} else {
					barcodeId = &decode.BarcodeId
					itemId = &decode.ItemId
				}
			}
			price := utils.StringToFloat64(info.OldPrice)
			priceAfterDisc := utils.StringToFloat64(info.MarketingPrice)
			disc := price - priceAfterDisc
			extCode := fmt.Sprintf("%d", *info.FboSku)
			daysOnSite := int32(time.Now().Sub(*info.CreatedAt).Hours() / 24)
			items := item.Items
			for i := 0; i < len(items); i++ {
				items[i].Price = &price
				items[i].PriceAfterDiscount = &priceAfterDisc
				items[i].Discount = &disc
				items[i].ExternalCode = &extCode
				items[i].DaysOnSite = &daysOnSite
				items[i].Barcode = info.Barcode
				items[i].BarcodeId = barcodeId
				items[i].ItemId = itemId
				items[i].Message = message
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *OzonService) getClient(host string) (*api.ClientWithResponses, error) {
	if c.client != nil {
		return c.client, nil
	}
	return api.NewClientWithResponses(c.mustHost(host), logging.WithStandardLoggerFn(), api.WithRequestEditorFn(c.customProvider))
}

func (c *OzonService) mustHost(host string) string {
	if len(host) == 0 {
		return "https://api-seller.ozon.ru"
	}
	return host
}

func (c *OzonService) prepareAttributes(resp *api.GetOzonProductAttributesResponse, transactionId int64) *[]model.StockItem {
	length := len(*resp.JSON200.Result)
	newItems := make([]model.StockItem, length)
	for index, item := range *resp.JSON200.Result {
		si := &model.StockItem{
			TransactionID:   transactionId,
			SupplierArticle: item.OfferId,
		}
		for _, attr := range *item.Attributes {
			if *attr.AttributeId == 85 && len(*attr.Values) > 0 {
				si.Brand = (*attr.Values)[0].Value
			}
			if *attr.AttributeId == 9461 && len(*attr.Values) > 0 {
				si.Subject = (*attr.Values)[0].Value
			}
			if *attr.AttributeId == 8229 && len(*attr.Values) > 0 {
				si.Category = (*attr.Values)[0].Value
			}
		}
		newItems[index] = *si
	}
	return &newItems
}

func (c *OzonService) preparePrices(resp *api.GetOzonProductInfoResponse, transactionId int64) error {
	if resp.StatusCode() != 200 || resp.JSON200 == nil {
		return errors.New(fmt.Sprintf("Ozon resp %d %s Body: %s", resp.StatusCode(), resp.Status(), string(resp.Body)))
	}
	length := len(*resp.JSON200.Result.Items)
	newItems := make([]model.StockItem, length)
	for index, info := range *resp.JSON200.Result.Items {
		c.productInfoCache[*info.FboSku] = info.Barcode
		price := utils.StringToFloat64(info.OldPrice)
		priceAfterDisc := utils.StringToFloat64(info.MarketingPrice)
		disc := price - priceAfterDisc
		extCode := fmt.Sprintf("%d", *info.FboSku)
		daysOnSite := int32(time.Now().Sub(*info.CreatedAt).Hours() / 24)
		si := &model.StockItem{
			Price:              &price,
			PriceAfterDiscount: &priceAfterDisc,
			Discount:           &disc,
			Barcode:            info.Barcode,
			TransactionID:      transactionId,
			ExternalCode:       &extCode,
			CardCreated:        *info.CreatedAt,
			DaysOnSite:         &daysOnSite,
		}
		newItems[index] = *si
	}
	return storage.UpdatePrices(&newItems)
}

func (c *OzonService) createItems(items *[]api.RowItem, dt time.Time, transactionId int64, source string) (*[]*model.StockItem, error) {
	newItems := make([]*model.StockItem, len(*items))
	for index, item := range *items {
		si, err := mapper.MapRowItem(&item, &dt)
		if err != nil {
			logrus.Errorf("Couldn't map a value at row %d (%s)", index, err)
			return nil, errors.Join(fmt.Errorf("Couldn't map a value at row %d sku %d", index, item.Sku), err)
		}
		si.OwnerCode = c.ownerCode
		si.Source = source
		si.TransactionID = transactionId
		newItems[index] = si
		c.products.addProduct(*item.Sku, si)
	}
	return &newItems, nil
}

func (c *OzonService) prepareItems(items *[]api.RowItem, dt time.Time, transactionId int64, source string) error {
	newItems := make([]*model.StockItem, len(*items))
	decoder := dictionary.GetItemDecoder()
	for index, item := range *items {
		si, err := mapper.MapRowItem(&item, &dt)
		if err != nil {
			logrus.Errorf("Couldn't map a value at row %d (%s)", index, err)
			return err
		}
		var barcodeId, itemId, message *string
		if si.Barcode != nil {
			decode, err := decoder.Decode(c.ownerCode, source, *si.Barcode)
			if err != nil {
				s := err.Error()
				message = &s
			} else {
				barcodeId = &decode.BarcodeId
				itemId = &decode.ItemId
			}
		}
		si.OwnerCode = c.ownerCode
		si.Source = source
		si.TransactionID = transactionId
		si.BarcodeId = barcodeId
		si.ItemId = itemId
		si.Message = message
		newItems[index] = si
		c.products.addProduct(*item.Sku, si)
	}
	return storage.SaveStocksInBatches(&newItems, c.config.BatchSize)
}

func createRequest(sinceDate, toDate time.Time, limit, offset int64) *api.GetOzonFBOJSONRequestBody {
	return &api.GetOzonFBOJSONRequestBody{
		Dir: "asc",
		Filter: api.FBOFilterFilter{
			Status: "",
			Since:  sinceDate,
			To:     toDate,
		},
		Translit: true,
		With: api.FBOFilterWith{
			AnalyticsData: true,
			FinancialData: true,
		},
		Limit:  limit,
		Offset: offset,
	}
}

func (c *OzonService) loadOrders(ctx context.Context, client *api.ClientWithResponses, transactionId int64, source string) error {
	toDate := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)
	sinceDate := toDate.AddDate(0, 0, -1)

	lastDate := toDate.AddDate(0, 0, -c.config.Order.LoadedDays)
	req := createRequest(sinceDate, toDate, int64(c.config.BatchSize), 0)

	limit := int64(c.config.BatchSize)
	offset := int64(0)
	affected := int64(0)
	for {
		req.Offset = offset
		req.Limit = limit
		req.Filter.To = toDate
		req.Filter.Since = sinceDate

		response, err := getFBOList(ctx, client, c, req)
		if err != nil {
			return errors2.Errorf("Error from GetFBO List: %s", err)
		}
		count, err := c.parseFBO(response, transactionId, source, c.ownerCode)
		affected += count
		if err != nil {
			return err
		}
		if count < limit {
			offset = 0
			toDate = sinceDate
			sinceDate = toDate.AddDate(0, 0, -1)
			logrus.Debugf("Loaded date %s and %s", toDate.Format("2006-01-02"), sinceDate.Format("2006-01-02"))
			if lastDate.After(sinceDate) {
				logrus.Debugf("Loaded %d orders end date %s", affected, sinceDate.Format("2006-01-02"))
				return nil
			}
		} else {
			offset += limit
		}
	}
}

func getFBOList(ctx context.Context, client *api.ClientWithResponses, c *OzonService, req *api.GetOzonFBOJSONRequestBody) (*api.GetOzonFBOResponse, error) {

	attemption := 3

	for {
		ctxCancel, _ := context.WithTimeout(ctx, time.Second*time.Duration(c.config.Timeout))
		response, err := client.GetOzonFBOWithResponse(ctxCancel, *req)
		attemption--
		if err == nil || attemption == 0 {
			return response, err
		}
		time.Sleep(5 * time.Second)
	}
}

func (c *OzonService) parseFBO(FBOResponse *api.GetOzonFBOResponse, transactionId int64, source string, ownerCode string) (int64, error) {
	if FBOResponse.StatusCode() != 200 {
		return 0, errors.New(fmt.Sprintf("Ozon resp code %d %s Body: %s", FBOResponse.StatusCode(), FBOResponse.Status(), string(FBOResponse.Body)))
	}
	if FBOResponse.JSON200 == nil {
		return 0, errors.New("Ozon resp is nil")
	}
	FBOItems := FBOResponse.JSON200.Result
	count := len(*FBOItems)
	logrus.Debugf("Loaded %d orders", count)
	if count == 0 {
		return 0, nil
	}
	var orders []model.Order
	decoder := dictionary.GetItemDecoder()

	searchFun := func(sku int64) *string {
		info, ok := c.products.get(sku)
		if !ok {
			return nil
		}
		return info.Barcode
	}

	for _, item := range *FBOItems {
		o, err := mapper.MapFBOToOrder(&item, transactionId, source, ownerCode, searchFun, decoder)
		if err != nil {
			return 0, fmt.Errorf("Couldn't map a value at row %s (%w)", item, err)
		}
		orders = append(orders, *o...)
	}
	if err := storage.SaveOrders(&orders); err != nil {
		return 0, err
	}
	return int64(count), nil
}
