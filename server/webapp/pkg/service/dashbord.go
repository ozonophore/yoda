package service

import (
	"github.com/yoda/webapp/pkg/api"
	"github.com/yoda/webapp/pkg/dao"
)

func GetSalesForWeek() (*api.SalesForWeek, error) {
	items, err := dao.GetWeekSales()
	if err != nil {
		return nil, err
	}
	var resultItems []api.SalesForWeekItem
	for _, item := range *items {
		resultItems = append(resultItems, api.SalesForWeekItem{
			OrderDate: item.OrderDate,
			Price:     item.Price,
		})
	}
	lastDate, err := dao.GetSalesDeleveredLastUpdate()
	if err != nil {
		return nil, err
	}
	return &api.SalesForWeek{
		Items:    &resultItems,
		UpdateAt: lastDate,
	}, nil
}
