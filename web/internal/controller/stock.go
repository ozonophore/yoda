package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
	"io"
	"net/http"
	"time"
)

type OrderService interface {
	PrepareAndReturnExcel(writer io.Writer, date time.Time) error
	GetOrders(date time.Time, filter string, source string, page int32, size int32) (*api.Orders, error)
	GetOrdersProduct(dateFrom time.Time, dateTo time.Time, filter *string, offset int32, limit int32, groupBy *string) (*api.OrderProducts, error)
	ExportOrderProductReport(writer http.ResponseWriter, dateFrom time.Time, dateTo time.Time, filter *string, groupBy *string) error
}

type SaleService interface {
	PrepareAndReturnExcel(writer io.Writer, year uint16, month uint8) error
	GetSale(year uint16, month uint8, page int32, size int32) (*api.Sales, error)
}

type AuthService interface {
	CreateToken(login *api.LoginInfo) (string, time.Time, error)
	GetProfile(auth string) (*api.Profile, error)
}

type IStockService interface {
	GetStocks(stockDate time.Time, limit, offset int) (*api.StocksFull, error)
}

type DictionaryService interface {
	GetPositions(offset int32, limit int32, source []string, filter *string) (*api.DictPositions, error)
	ExportWarehouses(writer http.ResponseWriter, source []string, code *string, cluster *string) error
	GetWarehouses(offset int32, limit int32, source []string, code *string, cluster *string) (*api.Warehouses, error)
}

type Controller struct {
	store        *storage.Storage
	orderService OrderService
	saleService  SaleService
	authService  AuthService
	dictService  DictionaryService
	stockService IStockService
}

func NewController(store *storage.Storage,
	orderService OrderService,
	saleService SaleService,
	authService AuthService,
	dictService DictionaryService,
	stockService IStockService) *Controller {
	return &Controller{
		store:        store,
		orderService: orderService,
		saleService:  saleService,
		authService:  authService,
		dictService:  dictService,
		stockService: stockService,
	}
}

func (c *Controller) GetStocks(ctx echo.Context, date types.Date) error {
	items, err := c.store.GetStocksByDate(date.Time)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return err
	}
	stocks := &api.Stocks{
		Count: int32(len(*items)),
		Items: mapStocks(items),
	}
	return ctx.JSON(http.StatusOK, stocks)
}

func (c *Controller) GetStocksWithPages(ctx echo.Context, params api.GetStocksWithPagesParams) error {
	stocks, err := c.stockService.GetStocks(params.Date.Time, params.Limit, params.Offset)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, stocks)
}

func mapStocks(items *[]storage.Stock) []api.Stock {
	count := len(*items)
	stocks := make([]api.Stock, count, count)
	for i, item := range *items {
		stocks[i] = api.Stock{
			Barcode:        item.Barcode,
			MarketplaceId:  item.MarketplaceId,
			OrganizationId: item.OrgId,
			Quantity:       item.Quantity,
			Organization:   item.OrgCode,
			Marketplace:    item.MarketplaceCode,
		}
	}
	return stocks
}
