package mapper

import (
	"fmt"
	"github.com/yoda/app/internal/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"strings"
	"time"
)

func MapSaleArray(sales *[]api.SalesItem, transactionId int64, source *string, ownerCode string, callback func(item *api.SalesItem)) []model.Sale {
	var result []model.Sale
	for _, s := range *sales {
		//callback(s.Odid)
		item := s
		callback(&item)
		result = append(result, *MapSale(s, transactionId, source, ownerCode))
	}
	return result
}

func MapSale(s api.SalesItem, transactionId int64, source *string, ownerCode string) *model.Sale {
	date, _ := time.Parse(time.DateOnly+"T"+time.TimeOnly, *s.Date)
	isStorno := utils.IntToBoolean(s.IsStorno)
	externalCode := fmt.Sprintf("%d", s.NmId)
	var warehouse *string
	if s.WarehouseName != nil {
		s := strings.ToUpper(*s.WarehouseName)
		warehouse = &s
	}
	return &model.Sale{
		TransactionID:     transactionId,
		TransactionDate:   time.Now(),
		Barcode:           s.Barcode,
		Source:            *source,
		OwnerCode:         ownerCode,
		LastChangeDate:    s.LastChangeDate.ToTime(),
		LastChangeTime:    s.LastChangeDate.ToTime(),
		SaleDate:          date,
		SaleTime:          date,
		SupplierArticle:   s.SupplierArticle,
		TechSize:          s.TechSize,
		TotalPrice:        s.TotalPrice,
		DiscountPercent:   utils.IntToInt32(s.DiscountPercent),
		IsSupply:          s.IsSupply,
		IsRealization:     s.IsRealization,
		PromoCodeDiscount: s.PromoCodeDiscount,
		WarehouseName:     warehouse,
		CountryName:       s.CountryName,
		OblastOkrugName:   s.OblastOkrugName,
		RegionName:        s.RegionName,
		IncomeID:          s.IncomeID,
		SaleID:            s.SaleID,
		Odid:              s.Odid,
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
