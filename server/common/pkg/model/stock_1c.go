package model

import "time"

type Stock struct {
	StockDate time.Time `gorm:"column:stock_date"`
	ItemID    string    `gorm:"column:item_id"`
	Quantity  float64   `gorm:"column:quantity"`
}

func (*Stock) TableName() string {
	return `dl."stock1c"`
}
