package service

import (
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/webapp/internal/api"
	"github.com/yoda/webapp/internal/dao"
	"time"
)

func GetOrders(limit, offset int, date time.Time, search *string) (*api.OrderItems, error) {
	items, err := dao.GetPageOrders(limit, offset, date, search)
	if err != nil {
		return nil, err
	}
	total := 0
	if len(*items) > 0 {
		total = (*items)[0].Total
	}
	if total == 0 {
		return &api.OrderItems{
			Limit:  limit,
			Offset: offset,
			Total:  total,
			Items:  []api.OrderItem{},
		}, nil
	}
	return &api.OrderItems{
		Limit:  limit,
		Offset: offset,
		Total:  total,
		Items:  *mapOrdersItems(items),
	}, nil
}

func mapOrdersItems(source *[]model.OrderPageItem) *[]api.OrderItem {
	var items []api.OrderItem
	for _, item := range *source {
		items = append(items, api.OrderItem{
			Id:                item.ID,
			Article:           item.SupplierArticle,
			Barcode:           item.Barcode,
			Quantity:          item.Quantity,
			Warehouse:         item.WarehouseName,
			Marketplace:       item.Source,
			Organisation:      item.OwnerCode,
			Name:              item.Name,
			CreatedAt:         item.CreatedDate,
			Price:             item.TotalPrice,
			PriceWithDiscount: item.PriceWithDiscount,
			Status:            item.Status,
		})
	}
	return &items
}
