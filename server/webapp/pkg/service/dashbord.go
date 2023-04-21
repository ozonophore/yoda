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

func GetTransactionsInfo() (*api.TransactionsInfo, error) {
	info, err := dao.GetTransactionInfo()
	if err != nil {
		return nil, err
	}
	return &api.TransactionsInfo{
		LastStart: info.LastStart,
		LastEnd:   info.LastEnd,
		Total:     info.Total,
		Success:   info.Completed,
	}, nil
}

func GetStocksInfo() (*[]api.StockInfoItem, error) {
	infos, err := dao.GetStockInfo()
	if err != nil {
		return nil, err
	}
	var items []api.StockInfoItem
	if infos == nil {
		return &items, nil
	}
	for _, info := range *infos {
		items = append(items, api.StockInfoItem{
			ZeroQty:   info.ZeroQty,
			StockDate: info.StockDate,
			Total:     info.Total,
		})
	}
	return &items, nil
}
