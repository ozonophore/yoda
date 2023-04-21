package model

import "time"

type WeekSales struct {
	OrderDate time.Time `gorm:"column:order_date" json:"order_date"`
	Price     float64   `gorm:"column:price" json:"price"`
}

type TransactionInfo struct {
	Total     int        `gorm:"column:total" json:"total"`
	Completed int        `gorm:"column:completed" json:"completed"`
	LastStart *time.Time `gorm:"column:last_start" json:"last_start"`
	LastEnd   *time.Time `gorm:"column:last_end" json:"last_end"`
}
