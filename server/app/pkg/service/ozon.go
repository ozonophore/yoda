package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"net/http"
	"strconv"
	"time"
)

type OzonService struct {
	clientId         string
	apiKey           string
	ownerCode        string
	productInfoCache map[int64]*api.ProductInfo
	config           *configuration.Config
	client           *api.ClientWithResponses
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
		productInfoCache: make(map[int64]*api.ProductInfo),
	}
}

func (c *OzonService) customProvider(ctx context.Context, req *http.Request) error {
	req.Header.Set("Client-Id", c.clientId)
	req.Header.Set("Api-Key", c.apiKey)
	return nil
}

func (c *OzonService) Parsing(context context.Context, transactionID int64) error {
	logrus.Info("Start parsing ozon")
	client, err := c.getClient(c.config.Ozon.Host)
	if err != nil {
		return err
	}
	offset := 0
	whType := api.GetOzonSupplierStocksJSONBodyWarehouseTypeALL
	source := types.SourceTypeOzon
	dt := time.Now()
	//----------------- LOAD STOCK -----------------
	for {
		resp, err := client.GetOzonSupplierStocksWithResponse(context, api.GetOzonSupplierStocksJSONRequestBody{
			Limit:         &c.config.BatchSize,
			Offset:        &offset,
			WarehouseType: &whType,
		})
		if err != nil {
			return err
		}
		items := resp.JSON200.Result.Rows
		if err = c.prepareItems(items, dt, transactionID, source); err != nil {
			return err
		}
		length := len(*items)
		if length < c.config.BatchSize {
			break
		}
		offset = length
	}
	logrus.Info("Take an information about prices")
	suppArt := repository.SelectUniqueStockItem(transactionID)

	err = CallbackBatch[string](suppArt, c.config.BatchSize, func(batch *[]string) error {
		req := api.GetOzonProductInfoJSONRequestBody{
			OfferId: batch,
		}
		resp, err := client.GetOzonProductInfoWithResponse(context, req)
		if err != nil {
			return err
		}
		err = c.preparePrices(resp, transactionID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	logrus.Info("Take an information about product")
	visibility := api.ProductAttributeFilterFilterVisibilityALL
	var lastId *string
	err = CallbackBatch[string](suppArt, c.config.BatchSize, func(batch *[]string) error {
		limit := len(*batch)
		request := api.ProductAttributeFilter{
			Filter: &api.ProductAttributeFilterFilter{
				OfferId:    batch,
				Visibility: &visibility,
			},
			LastId: lastId,
			Limit:  &limit,
		}
		resp, err := client.GetOzonProductAttributesWithResponse(context, request)
		if err != nil {
			return err
		}
		if resp.StatusCode() == 404 {
			return errors.New(fmt.Sprintf("Ozon resp 404: %s", *resp.JSON404.Message))
		}
		lastId = resp.JSON200.LastId
		newItems := c.prepareAttributes(resp, transactionID)
		if err = repository.UpdateAttributes(newItems); err != nil {
			return err
		}
		return nil
	})
	//------------------ LOAD ORDERS ------------------
	logrus.Info("Load orders")
	if err := c.loadOrders(context, client, transactionID, source); err != nil {
		return err
	}
	logrus.Info("Finish to load orders")
	logrus.Info("Finish parsing ozon")
	return nil
}

func (c *OzonService) getClient(host string) (*api.ClientWithResponses, error) {
	if c.client != nil {
		return c.client, nil
	}
	return api.NewClientWithResponses(c.mustHost(host), api.WithRequestEditorFn(c.customProvider))
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
		return errors.New(fmt.Sprintf("Ozon resp %d", resp.StatusCode()))
	}
	length := len(*resp.JSON200.Result.Items)
	newItems := make([]model.StockItem, length)
	for index, info := range *resp.JSON200.Result.Items {
		c.productInfoCache[*info.FboSku] = &info
		price, err := strconv.ParseFloat(*info.OldPrice, 64)
		if err != nil {
			return err
		}
		priceAfterDisc, err := strconv.ParseFloat(*info.MarketingPrice, 64)
		if err != nil {
			return err
		}
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
	return repository.UpdatePrices(&newItems)
}

func (c *OzonService) prepareItems(items *[]api.RowItem, dt time.Time, transactionId int64, source string) error {
	newItems := make([]model.StockItem, len(*items))
	for index, item := range *items {
		si, err := mapper.MapRowItem(&item, &dt)
		if err != nil {
			logrus.Errorf("Couldn't map a value at row %d (%s)", index, err)
			return err
		}
		si.OwnerCode = c.ownerCode
		si.Source = source
		si.TransactionID = transactionId
		newItems[index] = *si
	}
	return repository.SaveStocks(&newItems)
}

func (c *OzonService) loadOrders(ctx context.Context, client *api.ClientWithResponses, transactionId int64, source string) error {
	toDate := time.Now()
	sinceDate := toDate.AddDate(0, 0, -c.config.Order.LoadedDays)
	return FetchBatch(ctx, int64(c.config.BatchSize), func(offset int64, limit int64) (int64, error) {
		filter := api.GetOzonFBOJSONRequestBody{
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
		ctxCancel, _ := context.WithTimeout(ctx, time.Second*time.Duration(c.config.Timeout))
		response, err := client.GetOzonFBOWithResponse(ctxCancel, filter)
		if err != nil {
			return 0, err
		}
		return c.parseFBO(response, transactionId, source, c.ownerCode)
	})
}

func (c *OzonService) parseFBO(FBOResponse *api.GetOzonFBOResponse, transactionId int64, source string, ownerCode string) (int64, error) {
	if FBOResponse.StatusCode() != 200 {
		return 0, errors.New(fmt.Sprintf("Ozon resp code %d", FBOResponse.StatusCode()))
	}
	if FBOResponse.JSON200 == nil {
		return 0, errors.New("Ozon resp is nil")
	}
	FBOItems := FBOResponse.JSON200.Result
	count := len(*FBOItems)
	if count == 0 {
		return 0, nil
	}
	var orders []model.Order
	for _, item := range *FBOItems {
		o := mapper.MapFBOToOrder(&item, transactionId, source, ownerCode, &c.productInfoCache)
		orders = append(orders, *o...)
	}
	if err := repository.SaveOrders(&orders); err != nil {
		return 0, err
	}
	return int64(count), nil
}
