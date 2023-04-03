package mapper

import (
	"fmt"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/common/pkg/utils"
	"testing"
	"time"
)

func TestMapSale(t *testing.T) {
	transactionId := int64(1)
	source := "source"
	barcode := "barcode"
	brand := "brand"
	category := "category"
	countryName := "countryName"
	date := "2020-11-01T12:30:02"
	discountPercent := 10
	finishedPrice := float32(10)
	forPay := float32(10)
	gNumber := "gNumber"
	incomeID := 56
	isRealization := true
	isStorno := 1
	isSupply := true
	lastChangeDate := "2020-11-01T12:30:02"
	nmId := 12
	oblastOkrugName := "oblastOkrugName"
	odid := 32
	priceWithDiscount := float32(2.5)
	promoCodeDiscount := float32(1.2)
	regionName := "regionName"
	saleID := "saleID"
	spp := float32(10)
	srid := "srid"
	sticker := "sticker"
	subject := "subject"
	supplierArticle := "supplierArticle"
	techSize := "techSize"
	totalPrice := float32(10.58)
	warehouseName := "warehouseName"
	s := api.SalesItem{
		Barcode:           &barcode,
		Brand:             &brand,
		Category:          &category,
		CountryName:       &countryName,
		Date:              &date,
		DiscountPercent:   &discountPercent,
		FinishedPrice:     &finishedPrice,
		ForPay:            &forPay,
		GNumber:           &gNumber,
		IncomeID:          &incomeID,
		IsRealization:     &isRealization,
		IsStorno:          &isStorno,
		IsSupply:          &isSupply,
		LastChangeDate:    types.StringToCustomTime(lastChangeDate),
		NmId:              &nmId,
		OblastOkrugName:   &oblastOkrugName,
		Odid:              &odid,
		PriceWithDisc:     &priceWithDiscount,
		PromoCodeDiscount: &promoCodeDiscount,
		RegionName:        &regionName,
		SaleID:            &saleID,
		Spp:               &spp,
		Srid:              &srid,
		Sticker:           &sticker,
		Subject:           &subject,
		SupplierArticle:   &supplierArticle,
		TechSize:          &techSize,
		TotalPrice:        &totalPrice,
		WarehouseName:     &warehouseName,
	}
	target := MapSale(s, transactionId, &source, "OWNERCODE")
	if target.TransactionID != transactionId {
		t.Errorf("MapSale() = %v, want %v", target.TransactionID, transactionId)
	}
	if target.Source != source {
		t.Errorf("MapSale() = %v, want %v", target.Source, source)
	}
	if *target.Barcode != *s.Barcode {
		t.Errorf("MapSale() = %v, want %v", target.Barcode, barcode)
	}
	if *target.Brand != *s.Brand {
		t.Errorf("MapSale() = %v, want %v", target.Brand, brand)
	}
	if *target.Category != *s.Category {
		t.Errorf("MapSale() = %v, want %v", target.Category, category)
	}
	if *target.CountryName != *s.CountryName {
		t.Errorf("MapSale() = %v, want %v", target.CountryName, countryName)
	}
	saleDateTime := target.SaleDate.Format(time.DateOnly) + "T" + target.SaleTime.Format(time.TimeOnly)
	if saleDateTime != *s.Date {
		t.Errorf("MapSale() = %v, want %v", saleDateTime, *s.Date)
	}
	if *target.DiscountPercent != *utils.IntToInt32(s.DiscountPercent) {
		t.Errorf("MapSale() = %v, want %v", target.DiscountPercent, *s.DiscountPercent)
	}
	if *target.FinishedPrice != *utils.Float32ToFloat64(s.FinishedPrice) {
		t.Errorf("MapSale() = %v, want %v", target.FinishedPrice, *s.FinishedPrice)
	}
	if *target.ForPay != *utils.Float32ToFloat64(s.ForPay) {
		t.Errorf("MapSale() = %v, want %v", target.ForPay, *s.ForPay)
	}
	if *target.GNumber != *s.GNumber {
		t.Errorf("MapSale() = %v, want %v", target.GNumber, *s.GNumber)
	}
	if *target.IncomeID != *utils.IntToInt32(s.IncomeID) {
		t.Errorf("MapSale() = %v, want %v", target.IncomeID, *s.IncomeID)
	}
	if *target.IsRealization != *s.IsRealization {
		t.Errorf("MapSale() = %v, want %v", target.IsRealization, *s.IsRealization)
	}
	if *target.IsStorno != utils.IntToBoolean(s.IsStorno) {
		t.Errorf("MapSale() = %v, want %v", target.IsStorno, *s.IsStorno)
	}
	if *target.IsSupply != *s.IsSupply {
		t.Errorf("MapSale() = %v, want %v", target.IsSupply, *s.IsSupply)
	}
	lastChangeDateTime := target.LastChangeDate.Format(time.DateOnly) + "T" + target.LastChangeTime.Format(time.TimeOnly)
	if lastChangeDateTime != s.LastChangeDate.ToString() {
		t.Errorf("MapSale() = %v, want %v", lastChangeDateTime, s.LastChangeDate.String())
	}
	if *target.ExternalCode != fmt.Sprintf("%d", s.NmId) {
		t.Errorf("MapSale() = %v, want %v", *target.ExternalCode, *s.NmId)
	}
	if *target.OblastOkrugName != *s.OblastOkrugName {
		t.Errorf("MapSale() = %v, want %v", target.OblastOkrugName, *s.OblastOkrugName)
	}
	if *target.Odid != *utils.IntToInt32(s.Odid) {
		t.Errorf("MapSale() = %v, want %v", target.Odid, *s.Odid)
	}
	if *target.PriceWithDisc != *utils.Float32ToFloat64(s.PriceWithDisc) {
		t.Errorf("MapSale() = %v, want %v", target.PriceWithDisc, *s.PriceWithDisc)
	}
	if *target.PromoCodeDiscount != *s.PromoCodeDiscount {
		t.Errorf("MapSale() = %v, want %v", *target.PromoCodeDiscount, *s.PromoCodeDiscount)
	}
	if *target.RegionName != *s.RegionName {
		t.Errorf("MapSale() = %v, want %v", target.RegionName, *s.RegionName)
	}
	if *target.SaleID != *s.SaleID {
		t.Errorf("MapSale() = %v, want %v", *target.SaleID, *s.SaleID)
	}
	if *target.Spp != *utils.Float32ToFloat64(s.Spp) {
		t.Errorf("MapSale() = %v, want %v", target.Spp, *s.Spp)
	}
	if *target.Srid != *s.Srid {
		t.Errorf("MapSale() = %v, want %v", *target.Srid, *s.Srid)
	}
	if *target.Sticker != *s.Sticker {
		t.Errorf("MapSale() = %v, want %v", target.Sticker, *s.Sticker)
	}
	if *target.Subject != *s.Subject {
		t.Errorf("MapSale() = %v, want %v", target.Subject, *s.Subject)
	}
	if *target.SupplierArticle != *s.SupplierArticle {
		t.Errorf("MapSale() = %v, want %v", target.SupplierArticle, *s.SupplierArticle)
	}
	if *target.TechSize != *s.TechSize {
		t.Errorf("MapSale() = %v, want %v", target.TechSize, *s.TechSize)
	}
	if *target.TotalPrice != *utils.Float32ToFloat64(s.TotalPrice) {
		t.Errorf("MapSale() = %v, want %v", target.TotalPrice, *s.TotalPrice)
	}
	if *target.WarehouseName != *s.WarehouseName {
		t.Errorf("MapSale() = %v, want %v", target.WarehouseName, *s.WarehouseName)
	}
}
