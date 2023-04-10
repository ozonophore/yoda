package model

import (
	"time"
)

const TableNameJob = "job"

// Job mapped from table <job>
type Job struct {
	ID          int                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateDate  *time.Time          `gorm:"column:create_date;default:now()" json:"create_date"`
	IsActive    bool                `gorm:"column:is_active;not null" json:"is_active"`
	Description *string             `gorm:"column:description" json:"description"`
	WeekDays    *string             `gorm:"column:week_days" json:"week_days"` // Дни недели monday | tuesday | wednesday | thursday | friday | saturday | sunday
	AtTime      *string             `gorm:"column:at_time" json:"at_time"`     // Время в формате 8:04;16:00
	Interval    *int                `gorm:"column:interval" json:"interval"`   // Интервал в cекундах
	MaxRuns     *int                `gorm:"column:max_runs" json:"max_runs"`   // Максимальное количество запусков
	Type        string              `gorm:"column:type;not null" json:"type"`  // Тип задачи: regular | interval
	NextRun     *time.Time          `gorm:"column:next_run" json:"next_run"`
	LastRun     *time.Time          `gorm:"column:last_run" json:"last_run"`
	Params      *[]OwnerMarketplace `gorm:"many2many:job_owner;foreignKey:ID;joinForeignKey:JobID;References:OwnerCode;joinReferences:OwnerCode" json:"owner_marketplace"`
}

// TableName Job's table name
func (*Job) TableName() string {
	return TableNameJob
}
