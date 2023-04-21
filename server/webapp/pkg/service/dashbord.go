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
	var result api.SalesForWeek
	for _, item := range *items {
		result = append(result, api.SalesForWeekItem{
			OrderDate: item.OrderDate,
			Price:     item.Price,
		})
	}
	return &result, nil
}
