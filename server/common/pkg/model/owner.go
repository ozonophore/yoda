package model

import (
	"time"
)

const TableNameOwner = "owner"

type Owner struct {
	Code       string    `gorm:"column:code;primaryKey" json:"code"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	CreateDate time.Time `gorm:"column:create_date;default:now()" json:"create_date"`
	IsDeleted  bool      `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

func (*Owner) TableName() string {
	return TableNameOwner
}
