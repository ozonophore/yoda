// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const (
	STATUS_CREATED = "CREATED"
	STATUS_RUNNING = "RUNNING"
	STATUS_PENDING = "PENDING"
	STATUS_STOPPED = "STOPPED"
)

const (
	SCHEEDULER_MAIN   = "MAIN"
	SCHEEDULER_SYSTEM = "SYSTEM"
)

const TableNameScheduler = `"ml"."scheduler"`

// Scheduler mapped from table <scheduler>
type Scheduler struct {
	Code        string    `gorm:"column:code;primaryKey" json:"code"`
	Description *string   `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status;not null" json:"status"`
	UpdateAt    time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

// TableName Scheduler's table name
func (*Scheduler) TableName() string {
	return TableNameScheduler
}
