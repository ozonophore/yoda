package controller

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
	"io"
	"net/http"
	"time"
)

type OrderService interface {
	PrepareAndReturnExcel(writer io.Writer, date time.Time) error
	GetOrders(date time.Time, filter string, source string, page int32, size int32) (*api.Orders, error)
}

type SaleService interface {
	PrepareAndReturnExcel(writer io.Writer, year uint16, month uint8) error
	GetSale(year uint16, month uint8, page int32, size int32) (*api.Sales, error)
}

type AuthService interface {
	CreateToken(login *api.LoginInfo) (string, time.Time, error)
	GetProfile(auth string) (*api.Profile, error)
}

type Controller struct {
	store        *storage.Storage
	orderService OrderService
	saleService  SaleService
	authService  AuthService
}

func NewController(store *storage.Storage, orderService OrderService, saleService SaleService, authService AuthService) *Controller {
	return &Controller{
		store:        store,
		orderService: orderService,
		saleService:  saleService,
		authService:  authService,
	}
}

func (c *Controller) GetStocksDate(ctx echo.Context, date openapi_types.Date) error {
	items, err := c.store.GetStocksByDate(date.Time)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return err
	}
	stocks := &api.Stocks{
		Count: int32(len(*items)),
		Items: mapStocks(items),
	}
	ctx.JSON(http.StatusOK, stocks)
	return nil
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
