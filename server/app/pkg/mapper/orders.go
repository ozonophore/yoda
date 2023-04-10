package mapper

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"time"
)

func MapOrderArray(orders *[]api.OrdersItem, transactionId int64, source string, ownerCode string) (*[]model.Order, error) {
	var result []model.Order
	for _, order := range *orders {
		newOrder, err := MapOrder(&order, transactionId, source, ownerCode)
		if err != nil {
			return nil, err
		}
		result = append(result, *newOrder)
	}
	return &result, nil
}

func MapOrder(order *api.OrdersItem, transactionId int64, source string, ownerCode string) (result *model.Order, err error) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		bytes, err := json.Marshal(order)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Debugf("Order: %s", string(bytes))
	}
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
	discountPercent := float64(*order.DiscountPercent)
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
		Odid:              utils.IntToInt32(order.Odid),
		Subject:           order.Subject,
		Category:          order.Category,
		Brand:             order.Brand,
		IsCancel:          utils.BooleanToBoolean(order.IsCancel),
		CancelDt:          cancelDt,
		GNumber:           order.GNumber,
		Sticker:           order.Sticker,
		Srid:              order.Srid,
		Quantity:          int64(1),
	}, nil
}
