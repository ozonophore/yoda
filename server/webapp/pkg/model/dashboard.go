package model

import "time"

type WeekSales struct {
	OrderDate time.Time `gorm:"column:order_date" json:"order_date"`
	Price     float64   `gorm:"column:price" json:"price"`
}
