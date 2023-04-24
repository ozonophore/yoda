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

func GetTasksInfo() (*api.TaskInfo, error) {
	infos, err := dao.GetTasksInfo()
	if err != nil {
		return nil, err
	}
	var items = make([]api.TaskInfoItem, 0)
	if (*infos)[0].Completed == 0 && (*infos)[0].Canceled == 0 {
		return &api.TaskInfo{
			Items:     &items,
			Completed: 0,
			Canceled:  0,
		}, nil
	}
	for _, info := range *infos {
		items = append(items, api.TaskInfoItem{
			Id:        info.ID,
			StartDate: info.StartDate,
			EndDate:   info.EndDate,
			Status:    info.Status,
			Message:   info.Message,
		})
	}
	return &api.TaskInfo{
		Items:     &items,
		Completed: (*infos)[0].Completed,
		Canceled:  (*infos)[0].Canceled,
	}, nil
}
