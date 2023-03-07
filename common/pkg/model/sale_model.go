package model

import "time"

type Sale struct {
	ID                int64 `gorm:"primaryKey"`
	TransactionId     *int64
	Source            *string
	LastChangeDate    time.Time
	LastChangeTime    time.Time
	Date              time.Time
	SupplierArticle   *string
	TechSize          *string
	Barcode           *string
	TotalPrice        *float64
	DiscountPercent   *int32
	IsSupply          *bool
	IsRealization     *bool
	PromoCodeDiscount *float32
	WarehouseName     *string
	CountryName       *string
	OblastOkrugName   *string
	RegionName        *string
	IncomeId          *int32
	SaleId            *string
	Odid              *int32
	Spp               *float32
	ForPay            *float32
	FinishedPrice     *float64
	PriceWithDisc     *float64
	NmId              *int32
	Subject           *string
	Category          *string
	Brand             *string
	IsStorno          *bool
	GNumber           *string
	Sticker           *string
	Srid              *string
}

func (Sale) TableName() string {
	return "sale"
}
