package mapper

import (
	"fmt"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"time"
)

func MapSale(s api.SalesItem, transactionId *int64, source *string) *model.Sale {
	changeDate, _ := time.Parse(time.DateOnly+"T"+time.TimeOnly, *s.LastChangeDate)
	date, _ := time.Parse(time.DateOnly+"T"+time.TimeOnly, *s.Date)
	isStorno := utils.IntToBoolean(s.IsStorno)
	externalCode := fmt.Sprintf("%d", s.NmId)
	return &model.Sale{
		TransactionID:     *transactionId,
		Barcode:           s.Barcode,
		Source:            *source,
		LastChangeDate:    changeDate,
		LastChangeTime:    changeDate,
		SaleDate:          date,
		SaleTime:          date,
		SupplierArticle:   s.SupplierArticle,
		TechSize:          s.TechSize,
		TotalPrice:        utils.Float32ToFloat64(s.TotalPrice),
		DiscountPercent:   utils.IntToInt32(s.DiscountPercent),
		IsSupply:          s.IsSupply,
		IsRealization:     s.IsRealization,
		PromoCodeDiscount: s.PromoCodeDiscount,
		WarehouseName:     s.WarehouseName,
		CountryName:       s.CountryName,
		OblastOkrugName:   s.OblastOkrugName,
		RegionName:        s.RegionName,
		IncomeID:          utils.IntToInt32(s.IncomeID),
		SaleID:            s.SaleID,
		Odid:              utils.IntToInt32(s.Odid),
		Spp:               utils.Float32ToFloat64(s.Spp),
		ForPay:            utils.Float32ToFloat64(s.ForPay),
		FinishedPrice:     utils.Float32ToFloat64(s.FinishedPrice),
		PriceWithDisc:     utils.Float32ToFloat64(s.PriceWithDisc),
		ExternalCode:      &externalCode,
		Subject:           s.Subject,
		Category:          s.Category,
		Brand:             s.Brand,
		IsStorno:          &isStorno,
		GNumber:           s.GNumber,
		Sticker:           s.Sticker,
		Srid:              s.Srid,
	}
}
