package model

import "time"

type StockPageItem struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:subject;not null" json:"subject"`
	TransactionDate string    `gorm:"column:transaction_date;not null" json:"transaction_date"`
	OwnerCode       string    `gorm:"column:owner_code;not null" json:"owner_code"`
	Source          string    `gorm:"column:source;not null" json:"source"`
	SupplierArticle string    `gorm:"column:supplier_article;not null" json:"supplier_article"`
	Barcode         string    `gorm:"column:barcode;not null" json:"barcode"`
	Quantity        int       `gorm:"column:quantity;not null" json:"quantity"`
	QuantityFull    int       `gorm:"column:quantity_full;not null" json:"quantity_full"`
	WarehouseName   string    `gorm:"column:warehouse_name;not null" json:"warehouse_name"`
	Total           int       `gorm:"column:total;not null" json:"total"`
	CardCreated     time.Time `gorm:"column:card_created;not null" json:"card_created"`
}
