package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"log"
	"net/http"
	"strconv"
	"time"
)

type OzonService struct {
	clientId  string
	apiKey    string
	ownerCode string
	host      string
	config    *configuration.Configuration
	client    *api.ClientWithResponses
}

func NewOzonService(ownerCode, clientId, apiKey, host string, config *configuration.Configuration) *OzonService {
	return NewOzonServiceWIthClient(ownerCode, clientId, apiKey, host, nil, config)
}

func NewOzonServiceWIthClient(ownerCode, clientId, apiKey, host string, client *api.ClientWithResponses, config *configuration.Configuration) *OzonService {
	return &OzonService{
		clientId:  clientId,
		apiKey:    apiKey,
		config:    config,
		host:      host,
		ownerCode: ownerCode,
		client:    client,
	}
}

func (c *OzonService) customProvider(ctx context.Context, req *http.Request) error {
	req.Header.Set("Client-Id", c.clientId)
	req.Header.Set("Api-Key", c.apiKey)
	return nil
}

func (c *OzonService) Parsing(repository *repository.RepositoryDAO, jobId *int) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	offset := 0
	whType := api.GetOzonSupplierStocksJSONBodyWarehouseTypeALL
	source := types.SourceTypeOzon
	transactionId := (*repository).BeginOperation(&c.ownerCode, &source, jobId)
	dt := time.Now()
	for {
		resp, err := client.GetOzonSupplierStocksWithResponse(context.Background(), api.GetOzonSupplierStocksJSONRequestBody{
			Limit:         &c.config.BatchSize,
			Offset:        &offset,
			WarehouseType: &whType,
		})
		if err != nil {
			return err
		}
		items := resp.JSON200.Result.Rows
		if err = c.prepareItems(repository, items, dt, transactionId, source); err != nil {
			return err
		}
		length := len(*items)
		if length < c.config.BatchSize {
			break
		}
		offset = length
	}
	log.Println("Take an information about prices")
	suppArt := (*repository).SelectUniqueStockItem(transactionId)

	err = CallbackBatch[string](suppArt, c.config.BatchSize, func(batch *[]string) error {
		req := api.GetOzonProductInfoJSONRequestBody{
			OfferId: batch,
		}
		resp, err := client.GetOzonProductInfoWithResponse(context.Background(), req)
		if err != nil {
			return err
		}
		err = c.preparePrices(repository, resp, transactionId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Println("Take an information about product")
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
		resp, err := client.GetOzonProductAttributesWithResponse(context.Background(), request)
		if err != nil {
			return err
		}
		if resp.StatusCode() == 404 {
			return errors.New(fmt.Sprintf("Ozon resp 404: %s", *resp.JSON404.Message))
		}
		lastId = resp.JSON200.LastId
		newItems := c.prepareAttributes(resp, transactionId)
		if err = (*repository).UpdateAttributes(newItems); err != nil {
			return err
		}
		return nil
	})
	(*repository).EndOperation(transactionId, types.StatusTypeCompleted)
	log.Printf("Transaction %d was complited", *transactionId)
	return nil
}

func (c *OzonService) getClient() (*api.ClientWithResponses, error) {
	if c.client != nil {
		return c.client, nil
	}
	return api.NewClientWithResponses(c.mustHost(c.host), api.WithRequestEditorFn(c.customProvider))
}

func (c *OzonService) mustHost(host string) string {
	if len(host) == 0 {
		return "https://api-seller.ozon.ru"
	}
	return host
}

func (c *OzonService) prepareAttributes(resp *api.GetOzonProductAttributesResponse, transactionId *int64) *[]model.StockItem {
	length := len(*resp.JSON200.Result)
	newItems := make([]model.StockItem, length)
	for index, item := range *resp.JSON200.Result {
		si := &model.StockItem{
			TransactionID:   *transactionId,
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

func (c *OzonService) preparePrices(repository *repository.RepositoryDAO, resp *api.GetOzonProductInfoResponse, transactionId *int64) error {
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
			TransactionID:      *transactionId,
			ExternalCode:       &extCode,
		}
		newItems[index] = *si
	}
	return (*repository).UpdatePrices(&newItems)
}

func (c *OzonService) prepareItems(repository *repository.RepositoryDAO, items *[]api.RowItem, dt time.Time, transactionId *int64, source string) error {
	newItems := make([]model.StockItem, len(*items))
	for index, item := range *items {
		si, err := mapper.MapRowItem(&item, &dt)
		if err != nil {
			(*repository).EndOperation(transactionId, types.StatusTypeRejected)
			log.Printf("Couldn't map a value at row %d (%s)", index, err)
			return err
		}
		si.Source = source
		si.TransactionID = *transactionId
		newItems[index] = *si
	}
	return repository.SaveStocks(&newItems)
}
