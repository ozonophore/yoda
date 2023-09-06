package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func (s *Repository) CallReportOrdersByDay(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	startTime := time.Now()
	logrus.Debugf("Start calc report orders by day %s", startTime)
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("call dl.calc_report_order_by_day(?)", id).Error
		return err
	})
	logrus.Debugf("End calc report orders by day %s", time.Since(startTime))
	if err != nil {
		return fmt.Errorf("call calc_report_order_by_day with date %s error: %w", id, err)
	}
	return nil
}
