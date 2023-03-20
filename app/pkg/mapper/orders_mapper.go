package mapper

import (
	"fmt"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"time"
)

func MapOrderArray(orders *[]api.OrdersItem, transactionId int64, source string) (*[]model.Order, error) {
	var result []model.Order
	for _, order := range *orders {
		newOrder, err := MapOrder(&order, transactionId, source)
		if err != nil {
			return nil, err
		}
		result = append(result, *newOrder)
	}
	return &result, nil
}

func MapOrder(order *api.OrdersItem, transactionId int64, source string) (result *model.Order, err error) {
	changeDate, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *order.LastChangeDate)
	if err != nil {
		return nil, err
	}
	orderDate, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *order.Date)
	if err != nil {
		return nil, err
	}
	var cancelDt *time.Time
	if order.CancelDt != nil {
		cancelDtP, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *order.CancelDt)
		if err != nil {
			return nil, err
		}
		cancelDt = &cancelDtP
	}
	return &model.Order{
		TransactionID:   transactionId,
		Source:          source,
		LastChangeDate:  changeDate,
		LastChangeTime:  changeDate,
		OrderDate:       orderDate,
		OrderTime:       orderDate,
		SupplierArticle: order.SupplierArticle,
		TechSize:        order.TechSize,
		Barcode:         order.Barcode,
		TotalPrice:      float64(*order.TotalPrice),
		DiscountPercent: order.DiscountPercent,
		WarehouseName:   order.WarehouseName,
		Oblast:          order.Oblast,
		IncomeID:        order.IncomeID,
		ExternalCode:    fmt.Sprintf(`%d`, *order.NmId),
		Odid:            utils.IntToInt32(order.Odid),
		Subject:         order.Subject,
		Category:        order.Category,
		Brand:           order.Brand,
		IsCancel:        utils.BooleanToBoolean(order.IsCancel),
		CancelDt:        cancelDt,
		GNumber:         order.GNumber,
		Sticker:         order.Sticker,
		Srid:            order.Srid,
	}, nil
}
