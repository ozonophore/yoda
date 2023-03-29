package model

import (
	"time"
)

const TableNameJob = "job"

// Job mapped from table <job>
type Job struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OwnerCode     string         `gorm:"column:owner_code;not null" json:"owner_code"`
	CreateDate    *time.Time     `gorm:"column:create_date;default:now()" json:"create_date"`
	IsActive      bool           `gorm:"column:is_active;not null" json:"is_active"`
	Description   *string        `gorm:"column:description" json:"description"`
	JobParameters []JobParameter `gorm:"foreignKey:JobID;references:ID" json:"job_parameters"`
	Owner         Owner          `gorm:"foreignKey:OwnerCode;references:Code" json:"owner"`
}

// TableName Job's table name
func (*Job) TableName() string {
	return TableNameJob
}
