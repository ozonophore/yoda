package model

import "time"

type StockItem struct {
	ID                 int64 `gorm:"primaryKey"`
	TransactionId      *int64
	LastChangeDate     time.Time
	LastChangeTime     time.Time
	Source             *string
	SupplierArticle    *string
	Barcode            *string
	WarehouseName      *string
	ExternalCode       *string
	Name               *string
	Subject            *string
	Category           *string
	Brand              *string
	Price              *float64
	Discount           *float64
	PriceAfterDiscount *float64
	Quantity           *int32
	QuantityFull       *int32
	QuantityPromised   *int32
	DaysOnSite         *int32
}

func (StockItem) TableName() string {
	return "stock"
}
