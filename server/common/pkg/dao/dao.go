package dao

import (
	"errors"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func NewDao(database *gorm.DB) {
	db = database
}

func CreateScheduler(code string) error {
	err := db.Save(&model.Scheduler{
		Code:     code,
		Status:   model.STATUS_CREATED,
		UpdateAt: time.Now(),
	}).Error
	return errors.Join(err)
}

func GetSchedulerByCode(code string) (*model.Scheduler, error) {
	var scheduler model.Scheduler
	err := db.Where("code = ?", code).First(&scheduler).Error
	return &scheduler, errors.Join(err)
}

func GetScheduler() (*[]model.Scheduler, error) {
	var scheduler []model.Scheduler
	err := db.Find(&scheduler).Error
	return &scheduler, errors.Join(err)
}

func UpdateScheduler(code, status string) error {
	err := db.Model(&model.Scheduler{}).Where("code = ?", code).Updates(model.Scheduler{
		Status:   status,
		UpdateAt: time.Now(),
	}).Error
	return errors.Join(err)
}
