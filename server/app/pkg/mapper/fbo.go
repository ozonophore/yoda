package mapper

import (
	"fmt"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/model"
	"time"
)

func MapFBOToOrder(fbo *api.FBO, transactionId int64, source string, ownerCode string) *[]model.Order {
	var orders = make([]model.Order, len(*fbo.Products))
	var finData = make(map[int64]*api.PostingFinancialDataProduct)

	for _, f := range *fbo.FinancialData.Products {
		finData[f.ProductId] = &f
	}
	gNumber := fmt.Sprintf(`%d`, *fbo.OrderId)

	for i, product := range *fbo.Products {
		orders[i] = model.Order{
			TransactionID:     transactionId,
			TransactionDate:   time.Now(),
			Source:            source,
			OwnerCode:         ownerCode,
			OrderDate:         fbo.CreatedAt,
			OrderTime:         fbo.CreatedAt,
			SupplierArticle:   &product.OfferId,
			Barcode:           nil,
			TotalPrice:        finData[product.Sku].OldPrice,
			DiscountPercent:   finData[product.Sku].TotalDiscountPercent,
			DiscountValue:     finData[product.Sku].TotalDiscountValue,
			PriceWithDiscount: finData[product.Sku].Price,
			WarehouseName:     fbo.AnalyticsData.WarehouseName,
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
		}
	}
	return &orders
}
