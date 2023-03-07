package model

import "time"

type Order struct {
	ID              int64 `gorm:"primaryKey"`
	TransactionId   *int64
	Source          *string
	LastChangeDate  time.Time
	LastChangeTime  time.Time
	Date            time.Time
	SupplierArticle *string
	TechSize        *string
	Barcode         *string
	TotalPrice      *float64
	DiscountPercent *int32
	WarehouseName   *string
	Oblast          *string
	IncomeId        *int32
	NmId            *int32
	Odid            *int32
	Subject         *string
	Category        *string
	Brand           *string
	IsCancel        *bool
	CancelDt        time.Time
	GNumber         *string
	Sticker         *string
	Srid            *string
}

func (Order) TableName() string {
	return "order"
}
