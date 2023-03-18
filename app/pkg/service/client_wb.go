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

func (c *WBService) Parsing(repository *repository.RepositoryDAO, jobId *int) error {
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", c.apiKey)
	clnt, err := api.NewClientWithResponses("http://localhost:1080/wb", api.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return err
	}

	df := api.DateFrom{
		Time: time.Now(),
	}

	dateFrom := api.GetSupplierStocksParams{DateFrom: df}
	resp, err := clnt.GetSupplierStocksWithResponse(context.Background(), &dateFrom)
	if err != nil {
		return err
	}
	if resp.StatusCode()/100 != 2 {
		return errors.New(fmt.Sprintf("Http status: %s", resp.Status()))
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
		return err
	}
	if err = (*repository).SaveStocks(&newItems); err != nil {
		(*repository).EndOperation(transactionId, types.StatusTypeRejected)
		return err
	}

	respSales, err := clnt.GetWBSalesWithResponse(context.Background(), &api.GetWBSalesParams{
		DateFrom: df,
	})
	if err != nil {
		return err
	}
	if respSales.StatusCode() != 200 {
		return fmt.Errorf("Sales response status : %d", respSales.StatusCode())
	}
	salesItems := respSales.JSON200

	if err = CallbackBatch[api.SalesItem](salesItems, c.config.BatchSize, func(items *[]api.SalesItem) error {
		newItems := mapper.MapSaleArray(items, transactionId, &source)
		if err := repository.SaveSales(&newItems); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	(*repository).EndOperation(transactionId, types.StatusTypeCompleted)
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
