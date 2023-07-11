package stock

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
	"net/http"
)

type Server struct {
	store *storage.Storage
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		store: store,
	}
}

func (c *Server) GetStocksDate(ctx echo.Context, date openapi_types.Date) error {
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
