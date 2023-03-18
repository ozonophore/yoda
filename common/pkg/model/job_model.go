package model

import (
	"time"
)

const TableNameJob = "job"

// Job mapped from table <job>
type Job struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OwnerCode   string     `gorm:"column:owner_code;not null" json:"owner_code"`
	CreateDate  *time.Time `gorm:"column:create_date;default:now()" json:"create_date"`
	IsActive    bool       `gorm:"column:is_active;not null" json:"is_active"`
	Description *string    `gorm:"column:description" json:"description"`
}

// TableName Job's table name
func (*Job) TableName() string {
	return TableNameJob
}
