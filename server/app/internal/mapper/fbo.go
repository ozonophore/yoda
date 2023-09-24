package mapper

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/common/pkg/model"
	"strconv"
	"time"
)

func MapFBOToOrder(fbo *api.FBO, transactionId int64, source string, ownerCode string, cache func(int64) *string, decoder *dictionary.ItemDecoder) (*[]model.Order, error) {
	var orders = make([]model.Order, len(*fbo.Products))
	var finData = make(map[int64]*api.PostingFinancialDataProduct)

	finDate := fbo.FinancialData
	if finDate == nil || finDate.Products == nil {
		logrus.Error("financial data is empty for order %d", *fbo.OrderId)
	} else {
		for _, f := range *finDate.Products {
			finData[f.ProductId] = &f
		}
	}
	gNumber := fmt.Sprintf(`%d`, *fbo.OrderId)

	for i, product := range *fbo.Products {
		item := cache(product.Sku)
		var barcode *string
		if item != nil {
			barcode = item
		}
		var barcodeId, itemId, message *string
		if barcode != nil {
			decod, err := decoder.Decode(ownerCode, source, *barcode)
			if err != nil {
				s := err.Error()
				message = &s
			} else {
				barcodeId = &decod.BarcodeId
				itemId = &decod.ItemId
			}
		}
		supplierArticle := product.OfferId

		oldPrice := float64(0)
		totalDiscountPercent := float64(0)
		totalDiscountValue := float64(0)
		price, err := strconv.ParseFloat(product.Price, 64)
		if err != nil {
			return nil, errors.Errorf("invalid price %s for order %d", product.Price, *fbo.OrderId)
		}
		if finData != nil {
			oldPrice = finData[product.Sku].OldPrice
			totalDiscountPercent = finData[product.Sku].TotalDiscountPercent
			totalDiscountValue = finData[product.Sku].TotalDiscountValue
		}

		orders[i] = model.Order{
			TransactionID:     transactionId,
			TransactionDate:   time.Now(),
			Source:            source,
			OwnerCode:         ownerCode,
			OrderDate:         fbo.CreatedAt,
			OrderTime:         fbo.CreatedAt,
			SupplierArticle:   &supplierArticle,
			Barcode:           barcode,
			TotalPrice:        oldPrice,
			DiscountPercent:   totalDiscountPercent,
			DiscountValue:     totalDiscountValue,
			PriceWithDiscount: price,
			WarehouseName:     ToUpper(fbo.AnalyticsData.WarehouseName),
			Oblast:            fbo.AnalyticsData.Region,
			Region:            fbo.AnalyticsData.City,
			IncomeID:          nil,
			ExternalCode:      fmt.Sprintf(`%d`, product.Sku),
			Odid:              nil,
			Subject:           product.Name,
			Category:          nil,
			Brand:             nil,
			IsCancel:          fbo.Status == "cancelled",
			Status:            string(fbo.Status),
			CancelDt:          nil,
			GNumber:           &gNumber,
			Sticker:           nil,
			Srid:              fbo.PostingNumber,
			Quantity:          product.Quantity,
			BarcodeId:         barcodeId,
			ItemId:            itemId,
			Message:           message,
		}
	}
	return &orders, nil
}
