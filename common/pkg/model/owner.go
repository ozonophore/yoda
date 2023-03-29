package model

import (
	"time"
)

const TableNameOwner = "owner"

type Owner struct {
	Code       string `gorm:"primaryKey"`
	Name       string
	CreateDate time.Time
}

func (*Owner) TableName() string {
	return TableNameOwner
}
