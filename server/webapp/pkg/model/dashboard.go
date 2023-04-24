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

type StockInfo struct {
	StockDate time.Time `gorm:"column:stock_date" json:"stock_date"`
	Total     int       `gorm:"column:total" json:"total"`
	ZeroQty   int       `gorm:"column:zero_qty" json:"zero_qty"`
}

type TaskInfo struct {
	ID        int64      `gorm:"column:id" json:"id"`
	StartDate time.Time  `gorm:"column:start_date" json:"start_date"`
	EndDate   *time.Time `gorm:"column:end_date" json:"end_date"`
	Status    string     `gorm:"column:status" json:"status"`
	Message   *string    `gorm:"column:message" json:"message"`
	Completed int        `gorm:"column:completed" json:"completed"`
	Canceled  int        `gorm:"column:canceled" json:"canceled"`
}
