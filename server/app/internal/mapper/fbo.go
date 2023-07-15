package mapper

import (
	"fmt"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/app/internal/service/dictionary"
	"github.com/yoda/common/pkg/model"
	"time"
)

func MapFBOToOrder(fbo *api.FBO, transactionId int64, source string, ownerCode string, cache func(int64) *string, decoder *dictionary.ItemDecoder) *[]model.Order {
	var orders = make([]model.Order, len(*fbo.Products))
	var finData = make(map[int64]*api.PostingFinancialDataProduct)

	for _, f := range *fbo.FinancialData.Products {
		finData[f.ProductId] = &f
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
		orders[i] = model.Order{
			TransactionID:     transactionId,
			TransactionDate:   time.Now(),
			Source:            source,
			OwnerCode:         ownerCode,
			OrderDate:         fbo.CreatedAt,
			OrderTime:         fbo.CreatedAt,
			SupplierArticle:   &supplierArticle,
			Barcode:           barcode,
			TotalPrice:        finData[product.Sku].OldPrice,
			DiscountPercent:   finData[product.Sku].TotalDiscountPercent,
			DiscountValue:     finData[product.Sku].TotalDiscountValue,
			PriceWithDiscount: finData[product.Sku].Price,
			WarehouseName:     ToUpper(fbo.AnalyticsData.WarehouseName),
			Oblast:            fbo.AnalyticsData.Region,
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
	return &orders
}
