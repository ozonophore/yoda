package service

import (
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/webapp/pkg/api"
	"github.com/yoda/webapp/pkg/dao"
	"time"
)

func GetStocks(limit, offset int, date time.Time, search *string) (*api.StockItems, error) {
	items, err := dao.GetPageStocks(limit, offset, date, search)
	if err != nil {
		return nil, err
	}
	total := 0
	if len(*items) > 0 {
		total = (*items)[0].Total
	}
	if total == 0 {
		return &api.StockItems{
			Limit:  limit,
			Offset: offset,
			Total:  total,
			Items:  []api.StockItem{},
		}, nil
	}
	return &api.StockItems{
		Limit:  limit,
		Offset: offset,
		Total:  total,
		Items:  *mapStockItems(items),
	}, nil
}

func mapStockItems(source *[]model.StockPageItem) *[]api.StockItem {
	var items []api.StockItem
	for _, item := range *source {
		items = append(items, api.StockItem{
			Id:           item.ID,
			Article:      item.SupplierArticle,
			Barcode:      item.Barcode,
			Quantity:     item.Quantity,
			QuantityFull: item.QuantityFull,
			Warehouse:    item.WarehouseName,
			Marketplace:  item.Source,
			Organization: item.OwnerCode,
			Name:         item.Name,
			CreatedAt:    item.CardCreated,
		})
	}
	return &items
}
