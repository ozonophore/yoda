package service

import (
	"github.com/sirupsen/logrus"
	"time"
)

type stockDAOInterface interface {
	CallDailyAggr(day time.Time) error
	CalcDef(day time.Time) error
	CalcReport(day time.Time) error
	CalcDefByClusters(day time.Time) error
	CalcReportByClusters(day time.Time) error
	SetNotification(msg, sender, mtype string) error
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
	return s.dao.CalcReport(day)
}

func (s *StockService) CalcDefByClusters(day time.Time) error {
	return s.dao.CalcDefByClusters(day)
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

func (s *StockService) SetNotification(msg, sender, mtype string) error {
	return s.dao.SetNotification(msg, sender, mtype)
}
