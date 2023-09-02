package service

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type stockDAOInterface interface {
	CallDailyAggr(day time.Time) error
	CalcDef(day time.Time) error
	CalcReport(day time.Time) error
	CalcDefByClusters(day time.Time) error
	CalcReportByClusters(day time.Time) error
	CalcDefByItem(day time.Time) error
	CalcReportByItem(day time.Time) error
	SetNotification(msg, sender, mtype string) error

	CalcDefByProduct(day time.Time) error
	CalcReportByProduct(day time.Time) error
}

type StockService struct {
	dao    stockDAOInterface
	logger *logrus.Logger
}

func NewStockService(dao stockDAOInterface, logger *logrus.Logger) *StockService {
	return &StockService{
		dao:    dao,
		logger: logger,
	}
}

func (s *StockService) CalcAggrByDay(day time.Time) error {
	return s.dao.CallDailyAggr(day)
}

func (s *StockService) CalcDef(day time.Time) error {
	return s.dao.CalcDef(day)
}

func (s *StockService) CalcReport(day time.Time) error {
	i := 0
	for {
		err := s.dao.CalcReport(day)
		if err == nil {
			return nil
		}
		s.logger.Errorf("attept %d of CalcReport failed: %s", i, err)
		i++
		if i > 3 {
			return err
		}
	}
}

func (s *StockService) CalcDefByClusters(day time.Time) error {
	return s.dao.CalcDefByClusters(day)
}

func (s *StockService) CalcDefByItem(day time.Time) error {
	return s.dao.CalcDefByItem(day)
}

func (s *StockService) CalcDefAndReportByProduct(day time.Time) error {
	err := s.dao.CalcDefByProduct(day)
	if err != nil {
		return err
	}
	attemption := 3
	for {
		err = s.dao.CalcReportByProduct(day)
		if err == nil {
			break
		}
		attemption--
		if attemption == 0 {
			break
		}
	}
	if err != nil {
		return errors.Errorf("Error after 3 attemptions: %s", err)
	}
	return nil
}

func (s *StockService) CalcReportByClusters(day time.Time) error {
	i := 0
	for {
		err := s.dao.CalcReportByClusters(day)
		if err == nil {
			return nil
		}
		s.logger.Errorf("attept %d of CalcReportByClusters failed: %s", i, err)
		i++
		if i > 3 {
			return err
		}
	}
}

func (s *StockService) CalcReportByItem(day time.Time) error {
	i := 0
	for {
		err := s.dao.CalcReportByItem(day)
		if err == nil {
			return nil
		}
		s.logger.Errorf("attept %d of CalcReportByItem failed: %s", i, err)
		i++
		if i > 3 {
			return err
		}
	}
}

func (s *StockService) SetNotification(msg, sender, mtype string) error {
	return s.dao.SetNotification(msg, sender, mtype)
}
