package mapper

import (
	"fmt"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"time"
)

func MapOrderArray(orders *[]api.OrdersItem, transactionId int64, source string, ownerCode string, wasSold func(odid *int64) bool) (*[]model.Order, error) {
	var result []model.Order
	decoder := dictionary.GetItemDecoder()
	for _, order := range *orders {
		newOrder, err := MapOrder(&order, transactionId, source, ownerCode, wasSold)
		var barcodeId, itemId, message *string
		if newOrder.Barcode != nil {
			decode, err := decoder.Decode(ownerCode, source, *newOrder.Barcode)
			if err != nil {
				s := err.Error()
				message = &s
			} else {
				barcodeId = &decode.BarcodeId
				itemId = &decode.ItemId
			}
		}
		newOrder.BarcodeId = barcodeId
		newOrder.ItemId = itemId
		newOrder.Message = message
		if err != nil {
			return nil, err
		}
		result = append(result, *newOrder)
	}
	return &result, nil
}

func MapOrder(order *api.OrdersItem, transactionId int64, source string, ownerCode string, wasSold func(odid *int64) bool) (result *model.Order, err error) {
	orderDate, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *order.Date)
	if err != nil {
		return nil, err
	}
	var cancelDt *time.Time
	if order.CancelDt != nil && "0001-01-01T00:00:00" != *order.CancelDt {
		cancelDtP, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *order.CancelDt)
		if err != nil {
			return nil, err
		}
		cancelDt = &cancelDtP
	}
	status := "awaiting_deliver"
	if wasSold(order.Odid) {
		status = "delivered"
	} else if utils.BooleanToBoolean(order.IsCancel) {
		status = "canceled"
	}
	discountPercent := float64(*order.DiscountPercent)
	newSrid := fmt.Sprintf(`%d`, *order.Odid)
	return &model.Order{
		TransactionID:     transactionId,
		TransactionDate:   time.Now(),
		Source:            source,
		OwnerCode:         ownerCode,
		LastChangeDate:    order.LastChangeDate.ToTime(),
		LastChangeTime:    order.LastChangeDate.ToTime(),
		OrderDate:         orderDate,
		OrderTime:         orderDate,
		SupplierArticle:   order.SupplierArticle,
		TechSize:          order.TechSize,
		Barcode:           order.Barcode,
		TotalPrice:        *order.TotalPrice,
		DiscountPercent:   discountPercent,
		PriceWithDiscount: *order.TotalPrice * (1 - discountPercent/100),
		WarehouseName:     order.WarehouseName,
		Oblast:            order.Oblast,
		IncomeID:          order.IncomeID,
		ExternalCode:      fmt.Sprintf(`%d`, *order.NmId),
		Odid:              order.Odid,
		Subject:           order.Subject,
		Category:          order.Category,
		Brand:             order.Brand,
		IsCancel:          utils.BooleanToBoolean(order.IsCancel),
		CancelDt:          cancelDt,
		Status:            status,
		GNumber:           order.GNumber,
		Sticker:           order.Sticker,
		Srid:              &newSrid,
		Quantity:          int64(1),
	}, nil
}
