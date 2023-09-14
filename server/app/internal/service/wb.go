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
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"sync"
	"time"
)

type WBService struct {
	ownerCode string
	apiKey    string
	config    configuration.Config
	salesSet  map[int64]bool
}

func NewWBService(ownerCode, apiKey string, config *configuration.Config) *WBService {
	return &WBService{
		ownerCode: ownerCode,
		apiKey:    apiKey,
		config:    *config,
	}
}

func (c *WBService) Parsing(context context.Context, transactionID int64) error {
	logrus.Info("Start parsing wb ", c.ownerCode)
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", c.apiKey)
	clnt, err := api.NewClientWithResponses(c.config.Wb.Host, logging.WithStandardLoggerFn(), api.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return err
	}

	dateFrom := types.CustomTime(time.Now().AddDate(0, 0, -1*c.config.Wb.RemainingDays).UTC())
	logrus.Debugf("Date from: %s", dateFrom.String())

	request := api.GetSupplierStocksParams{DateFrom: dateFrom}
	resp, err := getSupplierStock(context, clnt, request)
	if err != nil {
		return err
	}
	items := *resp.JSON200
	length := len(items)
	if length == 0 {
		logrus.Warn("There aren't stocks")
	}
	source := types.SourceTypeWB
	ws := sync.WaitGroup{}
	ws.Add(1)
	go func() {
		defer ws.Done()
		c.extractReportDetailByPeriod(clnt, transactionID, source)
	}()

	newItems, err := c.prepareStocks(length, items, transactionID, source, c.ownerCode)
	if err != nil {
		return errors.Join(err)
	}
	if err = storage.SaveStocksInBatches(newItems, c.config.BatchSize); err != nil {
		return err
	}

	err = c.extractOrdersAndSales(context, transactionID, clnt, source)
	if err != nil {
		return fmt.Errorf("extractOrdersAndSales: %w", err)
	}
	ws.Wait()
	logrus.Info("Finish load orders from wb")
	logrus.Info("Finish parsing wb")
	return nil
}

func getSupplierStock(context context.Context, clnt *api.ClientWithResponses, request api.GetSupplierStocksParams) (*api.GetSupplierStocksResponse, error) {
	attempt := 3
	for {
		resp, err := clnt.GetSupplierStocksWithResponse(context, &request)
		if err == nil && resp.StatusCode() == 200 {
			return resp, err
		}
		attempt--
		if attempt == 0 {
			if err != nil {
				return nil, errors.New(fmt.Errorf("GetSupplierStocksWithResponse: %w", err).Error())
			}
			if resp.StatusCode()/100 != 2 {
				return nil, errors.New(fmt.Sprintf("Http status: %s msg: %s", resp.Status(), resp.Status()))
			}
		}
		time.Sleep(25 * time.Second)
	}
}

func (c *WBService) extractOrdersAndSales(ctx context.Context, transactionID int64, clnt *api.ClientWithResponses, source string) error {

	startDate := time.Now().AddDate(0, 0, 1).UTC()
	sinceDate := startDate.AddDate(0, 0, -1*c.config.Order.LoadedDays).UTC()
	logrus.Debugf("Sales since date: %s", sinceDate.String())
	count, err := c.extractSales(ctx, transactionID, clnt, sinceDate, source)
	logrus.Debugf("Sales count: %d", count)
	if err != nil {
		return err
	}

	startDate = time.Now().AddDate(0, 0, 1).UTC()
	sinceDate = startDate.AddDate(0, 0, -1*c.config.Order.LoadedDays).UTC()
	for {
		logrus.Debugf("Orders since date: %s", sinceDate.String())
		orders, err := c.callOrders(clnt, sinceDate)
		if err != nil {
			return err
		}
		start := time.Now()
		newOrders, lastChangeDate, err := mapper.MapOrderArray(orders, transactionID, source, c.ownerCode, func(item *int64) bool {
			return c.salesSet[*item]
		})
		if err != nil {
			return err
		}
		elapsed := time.Since(start)
		logrus.Debugf("Fetch took %s", elapsed)
		count, err := storage.SaveOrdersInBatches(newOrders, c.config.BatchSize)
		if err != nil {
			return err
		}
		logrus.Debugf("Orders count: %d", count)
		if len(*orders) < 70000 {
			break
		}
		sinceDate = lastChangeDate.Add(time.Second)
	}
	return nil
}

func (c *WBService) extractSales(ctx context.Context, transactionID int64, clnt *api.ClientWithResponses, df time.Time, source string) (int, error) {
	flag := 0
	dateFrom := types.CustomTime(df)
	i := 0
	var respSales *api.GetWBSalesResponse
	var err error
	for {
		ctxt, _ := context.WithTimeout(ctx, time.Duration(c.config.Timeout)*time.Second)
		respSales, err = clnt.GetWBSalesWithResponse(ctxt, &api.GetWBSalesParams{
			DateFrom: dateFrom,
			Flag:     &flag,
		})
		if err != nil || respSales.StatusCode() != 200 {
			i++
		} else {
			break
		}
		if i > 3 {
			break
		}
	}
	if err != nil {
		logrus.Debugf("Sales error: %s", err.Error())
		return 0, err
	}
	if respSales.StatusCode() != 200 {
		return 0, fmt.Errorf("Sales response status : %d", respSales.StatusCode())
	}
	salesItems := respSales.JSON200

	logrus.Debugf("Sales count: %d", len(*salesItems))
	if c.salesSet == nil {
		c.salesSet = make(map[int64]bool, len(*salesItems))
	}

	if err = CallbackBatch[api.SalesItem](salesItems, c.config.BatchSize, func(items *[]api.SalesItem) error {
		newItems := mapper.MapSaleArray(items, transactionID, &source, c.ownerCode, func(item *int64) {
			c.salesSet[*item] = true
		})
		if err := storage.SaveSales(&newItems); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return len(*salesItems), nil
}

func (c *WBService) callOrders(client *api.ClientWithResponses, df time.Time) (*[]api.OrdersItem, error) {
	flag := 0
	var dateFrom = types.CustomTime(df)
	request := api.GetWBOrdersParams{
		DateFrom: dateFrom,
		Flag:     &flag,
	}
	logrus.Debugf("Request to wb by date: %s", dateFrom.String())
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.config.Timeout)*time.Second)
	attemption := 3
	sleepTime := 30 * time.Second
	for {
		response, err := client.GetWBOrdersWithResponse(ctx, &request)
		if err == nil && response.StatusCode() == 200 {
			orders := response.JSON200
			return orders, err
		}
		attemption--
		if attemption == 0 {
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Response status : %d msg: %s", response.StatusCode(), err.Error()))
			}
			return nil, errors.New(fmt.Sprintf("Response status : %d msg: %s", response.StatusCode(), response.Status()))
		}
		time.Sleep(sleepTime)
	}
}

func (c *WBService) prepareStocks(length int, items []api.StocksItem, transactionId int64, source string, ownerCode string) (*[]*model.StockItem, error) {
	newItems := make([]*model.StockItem, length)
	decoder := dictionary.GetItemDecoder()
	for index, item := range items {
		var barcodeId, itemId, message *string
		if item.Barcode != nil {
			decode, err := decoder.Decode(ownerCode, source, *item.Barcode)
			if err != nil {
				s := err.Error()
				message = &s
			} else {
				barcodeId = &decode.BarcodeId
				itemId = &decode.ItemId
			}
		}

		si, err := mapper.MapStockItem(&item)
		si.OwnerCode = ownerCode
		si.BarcodeId = barcodeId
		si.ItemId = itemId
		si.Message = message
		priceAfterDiscount := *si.Price - *si.Discount
		si.PriceAfterDiscount = &priceAfterDiscount
		if err != nil {
			return nil, err
		}
		si.Source = source
		si.TransactionID = transactionId
		newItems[index] = si
	}
	return &newItems, nil
}

func (c *WBService) extractReportDetailByPeriod(client *api.ClientWithResponses, transactionId int64, source string) error {
	dateTo := types.CustomTime(time.Now())
	dateFrom := types.CustomTime(dateTo.ToTime().AddDate(0, 0, -7))

	resp, err := client.GetWBReportDetailByPeriodWithResponse(context.Background(), &api.GetWBReportDetailByPeriodParams{
		DateFrom: dateFrom,
		DateTo:   dateTo,
	})

	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("ReportDetailByPeriod response status : %d", resp.StatusCode())
	}
	items := resp.JSON200
	newItems := make([]model.ReportDetailByPeriod, len(*items))
	decoder := dictionary.GetItemDecoder()
	for index, item := range *items {
		var barcodeId, itemId, message *string
		if item.Barcode != nil {
			decode, err := decoder.Decode(c.ownerCode, source, *item.Barcode)
			if err != nil {
				s := err.Error()
				message = &s
			} else {
				barcodeId = &decode.BarcodeId
				itemId = &decode.ItemId
			}
		}
		newItems[index] = *mapper.MapReportDetailByPeriodItem(&item, transactionId, source, c.ownerCode, barcodeId, itemId, message)
	}
	err = storage.SaveReportDetail(&newItems, c.config.BatchSize)
	if err != nil {
		logrus.Error(err.Error())
	}
	return err
}
