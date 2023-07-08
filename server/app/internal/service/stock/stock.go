package service

import "time"

type stockDAOInterface interface {
	CallDailyAggr(day time.Time) error
	CalcDef(day time.Time) error
	CalcReport(day time.Time) error
	CalcDefByClusters(day time.Time) error
	CalcReportByClusters(day time.Time) error
	SetNotification(msg, sender, mtype string) error
}

type StockService struct {
	dao stockDAOInterface
}

func NewStockService(dao stockDAOInterface) *StockService {
	return &StockService{
		dao: dao,
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
	return s.dao.CalcReportByClusters(day)
}

func (s *StockService) SetNotification(msg, sender, mtype string) error {
	return s.dao.SetNotification(msg, sender, mtype)
}
