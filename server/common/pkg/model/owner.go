package model

import (
	"time"
)

const TableNameOwner = `"ml"."owner"`

type Owner struct {
	Code             string    `gorm:"column:code;primaryKey" json:"code"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	OrganisationId   *string   `gorm:"column:organisation_id" json:"organisation_id"`
	OrganisationName *string   `gorm:"column:organisation_name" json:"organisation_name"`
	CreateDate       time.Time `gorm:"column:create_date;default:now()" json:"create_date"`
	IsDeleted        bool      `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

func (*Owner) TableName() string {
	return TableNameOwner
}
