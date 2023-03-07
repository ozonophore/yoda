package model

import (
	"time"
)

type Job struct {
	ID          int64 `gorm:"primaryKey"`
	OwnerCode   *string
	Owner       Owner `gorm:"foreignKey:OwnerCode;references:OwnerCode"`
	CreateDate  time.Time
	IsActive    bool
	Description *string
}
