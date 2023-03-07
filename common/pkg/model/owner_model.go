package model

import (
	"time"
)

type Owner struct {
	Code       string `gorm:"primaryKey"`
	Name       string
	CreateDate time.Time
}
