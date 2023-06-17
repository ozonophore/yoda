package storage

import (
	"fmt"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) CallDailyAggr(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_daily_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_daily_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcDefFor30(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_def30_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_daily_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) GetJob(id int) (*model.Job, error) {
	var job model.Job
	err := s.db.Where(`"id"=?`, id).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (s *Repository) GetUniqueId() int64 {
	var id int64
	s.db.Raw(`select nextval('ml.log_id_seq')`).Scan(&id)
	return id
}