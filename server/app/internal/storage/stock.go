package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db     *gorm.DB
	config *configuration.Config
}

func NewRepository(db *gorm.DB, config *configuration.Config) *Repository {
	return &Repository{
		db:     db,
		config: config,
	}
}

func (s *Repository) CallDailyAggr(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_daily_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_daily_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcDef(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_def_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_def_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcDefByClusters(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_cluster_def_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_cluster_def_by_day with date %s error: %w", day, err)
	}
	return nil
}

// Дефектура по продукту
func (s *Repository) CalcDefByProduct(day time.Time) error {
	err := s.db.Exec("call dl.calc_item_def_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_item_def_by_day with date %s error: %w", day, err)
	}
	return nil
}

// Отчет по продутку
func (s *Repository) CalcReportByProduct(day time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	startTime := time.Now()
	logrus.Debugf("Start calc report(calc_report_by_product) %s", startTime)
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("call dl.calc_report_by_product(?)", day).Error
		return err
	})
	logrus.Debugf("End calc report(calc_report_by_product) %s", time.Since(startTime))
	if err != nil {
		return fmt.Errorf("call calc_report_by_product with date %s error: %w", day, err)
	}
	return nil
}

// Дефектура по кабинетам
func (s *Repository) CalcDefByOwner(day time.Time) error {
	err := s.db.Exec("call dl.calc_item_owner_def_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_item_def_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcDefByItem(day time.Time) error {
	err := s.db.Exec("call dl.calc_stock_item_def_by_day(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_stock_item_def_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcReport(day time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	startTime := time.Now()
	logrus.Debugf("Start calc report %s", startTime)
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("call dl.calc_sales_stock_by_day(?)", day).Error
		return err
	})
	logrus.Debugf("End calc report %s", time.Since(startTime))
	if err != nil {
		return fmt.Errorf("call calc_sales_stock_by_day with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcReportByItem(day time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	startTime := time.Now()
	logrus.Debugf("Start calc report %s", startTime)
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("call dl.calc_report_by_item(?)", day).Error
		return err
	})
	logrus.Debugf("End calc report by Item %s", time.Since(startTime))
	if err != nil {
		return fmt.Errorf("call calc_report_by_item with date %s error: %w", day, err)
	}
	return nil
}

func (s *Repository) CalcReportByClusters(day time.Time) error {
	err := s.db.Exec("call dl.calc_report_by_cluster(?)", day).Error
	if err != nil {
		return fmt.Errorf("call calc_report_by_cluster with date %s error: %w", day, err)
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

func (s *Repository) SetNotification(msg, sender, mtype string) error {
	return s.db.Exec(`insert into ml.notification (message, sender, type, is_sent) values (?, ?, ?, ?)`, msg, sender, mtype, false).Error
}
