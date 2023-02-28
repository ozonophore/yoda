package model

import "time"

type Table interface {
	TableName() string
}

type StockItem struct {
	ID              int64 `gorm:"primaryKey"`
	Transaction     *int64
	LastChangeDate  time.Time
	LastChangeTime  time.Time
	Source          *string
	SupplierArticle *string
	Barcode         *string
	WarehouseName   *string
	ExternalCode    *string
	Subject         *string
	Category        *string
	Brand           *string
	Price           *float64
	Discount        *float64
	Quantity        *int32
	QuantityFull    *int32
	DaysOnSite      *int32
}

func (StockItem) TableName() string {
	return "Stock"
}

type Transaction struct {
	ID        *int64     `gorm:"primaryKey"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time
	Type      *string `gorm:"not null"`
	Status    *string `gorm:"not null"`
}

func (Transaction) TableName() string {
	return "Transaction"
}
