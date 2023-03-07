package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/yoda/app/pkg/db"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/wbclient"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"log"
	"time"
)

type WBService struct {
	ownerCode string
	apiKey    string
}

func NewWBService(ownerCode string, apiKey string) *WBService {
	return &WBService{
		ownerCode: ownerCode,
		apiKey:    apiKey,
	}
}

func (c *WBService) Parsing(repository *db.RepositoryDAO, jobId *int) error {
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", c.apiKey)
	clnt, err := wbclient.NewClientWithResponses("http://localhost:1080/wb", wbclient.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return err
	}
	dateFrom := wbclient.GetSupplierStocksParams{DateFrom: wbclient.DateFrom{
		Time: time.Now(),
	}}
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
	newItems, err := c.prepareItems(repository, length, items, transactionId, source)
	if err != nil {
		(*repository).EndOperation(transactionId, types.StatusTypeRejected)
		return err
	}
	if err = (*repository).SaveStocks(&newItems); err != nil {
		(*repository).EndOperation(transactionId, types.StatusTypeRejected)
		return err
	}
	(*repository).EndOperation(transactionId, types.StatusTypeCompleted)
	return nil
}

func (c *WBService) prepareItems(repository *db.RepositoryDAO, length int, items []wbclient.StocksItem, transactionId *int64, source string) ([]model.StockItem, error) {
	newItems := make([]model.StockItem, length)
	for index, item := range items {
		si, err := mapper.MapStockItem(&item)
		priceAfterDiscount := *si.Price - *si.Discount
		si.PriceAfterDiscount = &priceAfterDiscount
		if err != nil {
			return nil, err
		}
		si.Source = &source
		si.TransactionId = transactionId
		newItems[index] = *si
	}
	return newItems, nil
}
