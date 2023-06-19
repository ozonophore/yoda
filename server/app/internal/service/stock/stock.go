package service

import "time"

type stockDAOInterface interface {
	CallDailyAggr(day time.Time) error
	CalcDef(day time.Time) error
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
