package mapper

import (
	"fmt"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/common/pkg/utils"
	"strings"
	"time"
)

func MapOrderArray(orders *[]api.OrdersItem, transactionId int64, source string, ownerCode string, soldMap map[int64]*model.DeliveredAddress) (*[]model.Order, *time.Time, error) {
	var result []model.Order
	var lastChangeDate *time.Time
	decoder := dictionary.GetItemDecoder()
	for _, order := range *orders {
		address, wasSold := soldMap[*order.Odid]
		newOrder, err := MapOrder(&order, transactionId, source, ownerCode, wasSold, address)
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
		if lastChangeDate == nil || newOrder.LastChangeDate.After(*lastChangeDate) {
			lastChangeDate = &newOrder.LastChangeDate
		}
		if err != nil {
			return nil, lastChangeDate, err
		}
		result = append(result, *newOrder)
	}
	return &result, lastChangeDate, nil
}

func ToUpper(v *string) *string {
	if v == nil {
		return nil
	}
	result := strings.ToUpper(*v)
	return &result
}

func MapOrder(order *api.OrdersItem, transactionId int64, source string, ownerCode string, wasSold bool, address *model.DeliveredAddress) (result *model.Order, err error) {
	orderDate := types.CustomTimeToTime(order.Date)
	if err != nil {
		return nil, err
	}
	var cancelDt *time.Time
	if order.CancelDt != nil && order.CancelDt.ToTime().Year() > 2000 {
		cancelDt = types.CustomTimeToTime(order.CancelDt)
	}
	status := "awaiting_deliver"
	if wasSold {
		status = "delivered"
	} else if utils.BooleanToBoolean(order.IsCancel) {
		status = "canceled"
	}
	discountPercent := float64(*order.DiscountPercent)
	newSrid := fmt.Sprintf(`%d`, *order.Odid)
	var (
		country *string
		region  *string
		oblast  *string
		saleDt  *time.Time
	)
	if address != nil {
		country = address.Country
		region = address.Region
		oblast = address.Okrug
		saleDt = &address.Date
	}

	return &model.Order{
		TransactionID:     transactionId,
		TransactionDate:   time.Now(),
		Source:            source,
		OwnerCode:         ownerCode,
		LastChangeDate:    order.LastChangeDate.ToTime(),
		LastChangeTime:    order.LastChangeDate.ToTime(),
		OrderDate:         *orderDate,
		OrderTime:         *orderDate,
		SupplierArticle:   order.SupplierArticle,
		TechSize:          order.TechSize,
		Barcode:           order.Barcode,
		TotalPrice:        *order.TotalPrice,
		DiscountPercent:   discountPercent,
		PriceWithDiscount: *order.TotalPrice * (1 - discountPercent/100),
		WarehouseName:     ToUpper(order.WarehouseName),
		Oblast:            oblast,
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
		Country:           country,
		Region:            region,
		SaleDate:          saleDt,
	}, nil
}
