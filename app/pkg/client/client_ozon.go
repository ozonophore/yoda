package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/db"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/wbclient"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type OzonService struct {
	clientId  string
	apiKey    string
	ownerCode string
	config    *configuration.Configuration
}

func NewOzonService(ownerCode, clientId, apiKey string, config *configuration.Configuration) *OzonService {
	return &OzonService{
		clientId:  clientId,
		apiKey:    apiKey,
		config:    config,
		ownerCode: ownerCode,
	}
}

func (c *OzonService) customProvider(ctx context.Context, req *http.Request) error {
	req.Header.Set("Client-Id", c.clientId)
	req.Header.Set("Api-Key", c.apiKey)
	return nil
}

func (c *OzonService) Parsing(listener *db.RepositoryDAO, jobId *int) error {
	clnt, err := wbclient.NewClientWithResponses("https://api-seller.ozon.ru", wbclient.WithRequestEditorFn(c.customProvider))
	if err != nil {
		return err
	}
	offset := 0
	whType := wbclient.GetOzonSupplierStocksJSONBodyWarehouseTypeALL
	source := types.SourceTypeOzon
	transactionId := (*listener).BeginOperation(&c.ownerCode, &source, jobId)
	dt := time.Now()
	for {
		resp, err := clnt.GetOzonSupplierStocksWithResponse(context.Background(), wbclient.GetOzonSupplierStocksJSONRequestBody{
			Limit:         c.config.BatchSize,
			Offset:        &offset,
			WarehouseType: &whType,
		})
		if err != nil {
			return err
		}
		items := resp.JSON200.Result.Rows
		if err = c.prepareItems(listener, items, dt, transactionId, source); err != nil {
			return err
		}
		length := len(*items)
		if length < *c.config.BatchSize {
			break
		}
		offset = length
	}
	log.Println("Take an information about prices")
	suppArt := (*listener).SelectUniqueStockItem(transactionId)
	var low int
	low = 0
	highest := len(*suppArt)
	for {
		high := low + int(math.Min(float64(*c.config.BatchSize), float64(highest-low)))
		batch := (*suppArt)[low:high]
		req := wbclient.GetOzonProductInfoJSONRequestBody{
			OfferId: &batch,
		}
		resp, err := clnt.GetOzonProductInfoWithResponse(context.Background(), req)
		if err != nil {
			return err
		}
		err = c.preparePrices(listener, resp, transactionId)
		if err != nil {
			return err
		}
		low += high
		if high == highest {
			break
		}
	}
	log.Println("Take an information about product")
	low = 0
	var lastId string
	for {
		limit := int(math.Min(float64(*c.config.BatchSize), float64(highest-low)))
		high := low + limit
		batch := (*suppArt)[low:high]
		visibility := wbclient.ProductAttributeFilterFilterVisibilityALL
		request := wbclient.ProductAttributeFilter{
			Filter: &wbclient.ProductAttributeFilterFilter{
				OfferId:    &batch,
				Visibility: &visibility,
			},
			LastId: &lastId,
			Limit:  &limit,
		}
		resp, err := clnt.GetOzonProductAttributesWithResponse(context.Background(), request)
		if err != nil {
			return err
		}
		if resp.StatusCode() == 404 {
			return errors.New(fmt.Sprintf("Ozon resp 404: $s", *resp.JSON404.Message))
		}
		lastId = *resp.JSON200.LastId
		newItems := c.prepareAttributes(resp, transactionId)
		if err = (*listener).UpdateAttributes(newItems); err != nil {
			return err
		}
		low += high
		if high == highest {
			break
		}
	}
	(*listener).EndOperation(transactionId, types.StatusTypeCompleted)
	log.Printf("Transaction %d was complited", *transactionId)
	return nil
}

func (c *OzonService) prepareAttributes(resp *wbclient.GetOzonProductAttributesResponse, transactionId *int64) *[]model.StockItem {
	length := len(*resp.JSON200.Result)
	newItems := make([]model.StockItem, length)
	for index, item := range *resp.JSON200.Result {
		si := &model.StockItem{
			TransactionId:   transactionId,
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

func (c *OzonService) preparePrices(repository *db.RepositoryDAO, resp *wbclient.GetOzonProductInfoResponse, transactionId *int64) error {
	length := len(*resp.JSON200.Result.Items)
	newItems := make([]model.StockItem, length)
	for index, info := range *resp.JSON200.Result.Items {
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
		si := &model.StockItem{
			Price:              &price,
			PriceAfterDiscount: &priceAfterDisc,
			Discount:           &disc,
			Barcode:            info.Barcode,
			TransactionId:      transactionId,
			ExternalCode:       &extCode,
		}
		newItems[index] = *si
	}
	return (*repository).UpdatePrices(&newItems)
}

func (c *OzonService) prepareItems(repository *db.RepositoryDAO, items *[]wbclient.RowItem, dt time.Time, transactionId *int64, source string) error {
	newItems := make([]model.StockItem, len(*items))
	for index, item := range *items {
		si, err := mapper.MapRowItem(&item, &dt)
		if err != nil {
			(*repository).EndOperation(transactionId, types.StatusTypeRejected)
			log.Printf("Couldn't map a value at row %d (%s)", index, err)
			return err
		}
		si.Source = &source
		si.TransactionId = transactionId
		newItems[index] = *si
	}
	return repository.SaveStocks(&newItems)
}
