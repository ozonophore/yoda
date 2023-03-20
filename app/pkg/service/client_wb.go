package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"log"
	"time"
)

type WBService struct {
	ownerCode string
	apiKey    string
	config    configuration.Configuration
}

func NewWBService(ownerCode string, apiKey string, config *configuration.Configuration) *WBService {
	return &WBService{
		ownerCode: ownerCode,
		apiKey:    apiKey,
		config:    *config,
	}
}

func (c *WBService) Parsing(repository *repository.RepositoryDAO, jobId *int) (*int64, error) {
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", c.apiKey)
	clnt, err := api.NewClientWithResponses("http://localhost:1080/wb", api.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return nil, err
	}

	dateFrom := api.DateFrom{
		Time: time.Now(),
	}

	request := api.GetSupplierStocksParams{DateFrom: dateFrom}
	resp, err := clnt.GetSupplierStocksWithResponse(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode()/100 != 2 {
		return nil, errors.New(fmt.Sprintf("Http status: %s", resp.Status()))
	}
	items := *resp.JSON200
	length := len(items)
	if length == 0 {
		log.Print("There aren't warehouses")
	}
	source := types.SourceTypeWB
	transactionId := (*repository).BeginOperation(&c.ownerCode, &source, jobId)

	newItems, err := c.prepareStocks(length, items, transactionId, source)
	if err != nil {
		(*repository).EndOperation(transactionId, types.StatusTypeRejected)
		return nil, err
	}
	if err = (*repository).SaveStocks(&newItems); err != nil {
		(*repository).EndOperation(transactionId, types.StatusTypeRejected)
		return nil, err
	}

	respSales, err := clnt.GetWBSalesWithResponse(context.Background(), &api.GetWBSalesParams{
		DateFrom: dateFrom,
	})
	if err != nil {
		return nil, err
	}
	if respSales.StatusCode() != 200 {
		return nil, fmt.Errorf("Sales response status : %d", respSales.StatusCode())
	}
	salesItems := respSales.JSON200

	if err = CallbackBatch[api.SalesItem](salesItems, c.config.BatchSize, func(items *[]api.SalesItem) error {
		newItems := mapper.MapSaleArray(items, transactionId, &source)
		if err := repository.SaveSales(&newItems); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	err = c.loadOrders(clnt, &dateFrom, repository, *transactionId, &source)
	if err != nil {
		return nil, err
	}
	(*repository).EndOperation(transactionId, types.StatusTypeCompleted)
	return transactionId, nil
}

/**
 * Load orders from WB
 */
func (c *WBService) loadOrders(client *api.ClientWithResponses, dateFrom *api.DateFrom, repository *repository.RepositoryDAO, transactionId int64, source *string) error {
	request := api.GetWBOrdersParams{
		DateFrom: *dateFrom,
	}
	response, err := client.GetWBOrdersWithResponse(context.Background(), &request)
	if err != nil {
		return err
	}
	orders := response.JSON200
	if err = CallbackBatch[api.OrdersItem](orders, c.config.BatchSize, func(items *[]api.OrdersItem) error {
		orders, errMap := mapper.MapOrderArray(items, transactionId, *source)
		if errMap != nil {
			return errMap
		}
		if err := repository.SaveOrders(orders); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *WBService) prepareStocks(length int, items []api.StocksItem, transactionId *int64, source string) ([]model.StockItem, error) {
	newItems := make([]model.StockItem, length)
	for index, item := range items {
		si, err := mapper.MapStockItem(&item)
		priceAfterDiscount := *si.Price - *si.Discount
		si.PriceAfterDiscount = &priceAfterDiscount
		if err != nil {
			return nil, err
		}
		si.Source = source
		si.TransactionID = *transactionId
		newItems[index] = *si
	}
	return newItems, nil
}
