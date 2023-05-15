package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/mapper"
	"github.com/yoda/app/internal/repository"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"time"
)

type WBService struct {
	ownerCode string
	apiKey    string
	config    configuration.Config
	salesOdid map[int64]bool
}

func NewWBService(ownerCode, apiKey string, config *configuration.Config) *WBService {
	return &WBService{
		ownerCode: ownerCode,
		apiKey:    apiKey,
		config:    *config,
		salesOdid: make(map[int64]bool),
	}
}

func (c *WBService) Parsing(context context.Context, transactionID int64) error {
	logrus.Info("Start parsing wb ", c.ownerCode)
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", c.apiKey)
	clnt, err := api.NewClientWithResponses(c.config.Wb.Host, logging.WithStandardLoggerFn(), api.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return err
	}

	dateFrom := api.DateFrom{
		Time: time.Now().AddDate(0, 0, -1*c.config.Wb.RemainingDays),
	}

	request := api.GetSupplierStocksParams{DateFrom: dateFrom}
	resp, err := clnt.GetSupplierStocksWithResponse(context, &request)
	if err != nil {
		return err
	}
	if resp.StatusCode()/100 != 2 {
		return errors.New(fmt.Sprintf("Http status: %s msg: %s", resp.Status(), resp.Status()))
	}
	items := *resp.JSON200
	length := len(items)
	if length == 0 {
		logrus.Warn("There aren't stocks")
	}
	source := types.SourceTypeWB
	newItems, err := c.prepareStocks(length, items, transactionID, source, c.ownerCode)
	if err != nil {
		return errors.Join(err)
	}
	if err = repository.SaveStocks(&newItems); err != nil {
		return err
	}

	//err = c.extractSales(context, transactionID, clnt, dateFrom, source)
	//if err != nil {
	//	return errors.Join(err)
	//}
	//logrus.Info("Load orders from wb")
	//err = c.extractOrders(clnt, transactionID, &source)
	//if err != nil {
	//	return errors.Join(err)
	//}
	c.extractOrdersAndSales(context, transactionID, clnt, dateFrom, source)
	logrus.Info("Finish load orders from wb")
	logrus.Info("Finish parsing wb")
	return nil
}

func (c *WBService) extractOrdersAndSales(ctx context.Context, transactionID int64, clnt *api.ClientWithResponses, dateFrom api.DateFrom, source string) error {
	e := make(chan error, 2)
	defer close(e)
	go func() {
		e <- c.extractSales(ctx, transactionID, clnt, dateFrom, source)
	}()
	o := make(chan *[]api.OrdersItem, 1)
	defer close(o)
	go func() {
		sinceDate := api.DateFrom{
			Time: time.Now(),
		}
		days := c.config.Order.LoadedDays
		sinceDate.Time = sinceDate.Time.AddDate(0, 0, -1*days)
		orders, err := c.callOrders(clnt, &sinceDate)
		e <- err
		o <- orders
	}()
	errs := []error{<-e, <-e}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	orders := <-o
	start := time.Now()
	newOrders, errMap := mapper.MapOrderArray(orders, transactionID, source, c.ownerCode, func(item *int64) bool {
		return c.salesOdid[*item]
	})
	elapsed := time.Since(start)
	logrus.Debugf("Fetch took %s", elapsed)
	if errMap != nil {
		return errMap
	}
	if err := repository.SaveOrdersInBatches(newOrders, c.config.BatchSize); err != nil {
		return err
	}
	return nil
}

func (c *WBService) extractSales(ctx context.Context, transactionID int64, clnt *api.ClientWithResponses, dateFrom api.DateFrom, source string) error {
	ctxt, _ := context.WithTimeout(ctx, time.Duration(c.config.Timeout)*time.Second)
	respSales, err := clnt.GetWBSalesWithResponse(ctxt, &api.GetWBSalesParams{
		DateFrom: dateFrom,
	})
	if err != nil {
		return err
	}
	if respSales.StatusCode() != 200 {
		return fmt.Errorf("Sales response status : %d", respSales.StatusCode())
	}
	salesItems := respSales.JSON200

	if err = CallbackBatch[api.SalesItem](salesItems, c.config.BatchSize, func(items *[]api.SalesItem) error {
		newItems := mapper.MapSaleArray(items, transactionID, &source, c.ownerCode, func(item *int64) {
			c.salesOdid[*item] = true
		})
		if err := repository.SaveSales(&newItems); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

/**
 * Load orders from WB
 */
func (c *WBService) extractOrders(client *api.ClientWithResponses, transactionId int64, source *string) error {
	sinceDate := api.DateFrom{
		Time: time.Now(),
	}
	days := c.config.Order.LoadedDays
	sinceDate.Time = sinceDate.Time.AddDate(0, 0, -1*days)
	if err := c.fetchOrders(client, &sinceDate, transactionId, source, c.ownerCode); err != nil {
		logrus.Debugf("Error while fetching orders from wb: %s", err.Error())
		return err
	}
	return nil
}

func (c *WBService) fetchOrders(client *api.ClientWithResponses, dateFrom *api.DateFrom, transactionId int64, source *string, ownerCode string) error {
	orders, err := c.callOrders(client, dateFrom)
	if err != nil {
		return err
	}
	if err = CallbackBatch[api.OrdersItem](orders, c.config.BatchSize, func(items *[]api.OrdersItem) error {
		start := time.Now()
		orders, errMap := mapper.MapOrderArray(items, transactionId, *source, ownerCode, func(item *int64) bool {
			return c.salesOdid[*item]
		})
		elapsed := time.Since(start)
		logrus.Debugf("Fetch took %s", elapsed)
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

func (c *WBService) callOrders(client *api.ClientWithResponses, dateFrom *api.DateFrom) (*[]api.OrdersItem, error) {
	flag := 0
	request := api.GetWBOrdersParams{
		DateFrom: *dateFrom,
		Flag:     &flag,
	}
	logrus.Debugf("Request to wb by date: %s", *dateFrom)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.config.Timeout)*time.Second)
	response, err := client.GetWBOrdersWithResponse(ctx, &request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("Response status : %d msg: %s", response.StatusCode(), response.Status()))
	}
	orders := response.JSON200
	return orders, err
}

func (c *WBService) prepareStocks(length int, items []api.StocksItem, transactionId int64, source string, ownerCode string) ([]model.StockItem, error) {
	newItems := make([]model.StockItem, length)
	for index, item := range items {
		si, err := mapper.MapStockItem(&item)
		si.OwnerCode = ownerCode
		priceAfterDiscount := *si.Price - *si.Discount
		si.PriceAfterDiscount = &priceAfterDiscount
		if err != nil {
			return nil, err
		}
		si.Source = source
		si.TransactionID = transactionId
		newItems[index] = *si
	}
	return newItems, nil
}
