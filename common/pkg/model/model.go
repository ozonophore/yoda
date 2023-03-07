package model

import "time"

type Table interface {
	TableName() string
}

type Transaction struct {
	ID        *int64     `gorm:"primaryKey"`
	OwnerCode *string    `gorm:"not null"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time
	Source    *string `gorm:"not null"`
	Status    *string `gorm:"not null"`
	JobId     *int    `gorm:"not null"`
}

func (Transaction) TableName() string {
	return "transaction"
}
