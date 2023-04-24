package dao

import (
	"github.com/yoda/webapp/pkg/model"
	"time"
)

func GetWeekSales() (*[]model.WeekSales, error) {
	var weekSales []model.WeekSales
	err := dao.database.Raw(`select order_date, price from (select order_date, sum(price_with_discount * quantity) price
               from order_delivered od
               where order_date > CURRENT_DATE - 10
                 and order_date <= CURRENT_DATE - 3
               group by order_date) data
order by 1`).Scan(&weekSales).Error
	if err != nil {
		return nil, err
	}
	return &weekSales, nil
}

type maxValue struct {
	Max *time.Time
}

func GetSalesDeleveredLastUpdate() (*time.Time, error) {
	var value maxValue
	err := dao.database.Raw(`select max(created_at) from order_delivered_log`).Scan(&value).Error
	if err != nil {
		return nil, err
	}
	return value.Max, nil
}

func GetTransactionInfo() (*model.TransactionInfo, error) {
	var jobGeneralInfo model.TransactionInfo
	err := dao.database.Raw(`with trs as (
    select max(id) id, count(1) cnt, sum( case when status = 'COMPLETED' then 1 else 0 end ) completed from "transaction"
	)
	select t.start_date last_start, t.end_date last_end, trs.cnt total, trs.completed from "transaction" t
         inner join trs on trs.id = t.id`).Scan(&jobGeneralInfo).Error
	if err != nil {
		return nil, err
	}
	return &jobGeneralInfo, nil
}

func GetStockInfo() (*[]model.StockInfo, error) {
	var stockInfo []model.StockInfo
	err := dao.database.Raw(`select stock_date, count(1) total, sum(case when quantity = 0 then 1 else 0 end) zero_qty from stock_daily
	where stock_date > current_date - 10
	group by stock_date
	order by stock_date`).Scan(&stockInfo).Error
	if err != nil {
		return nil, err
	}
	return &stockInfo, nil
}

func GetTasksInfo() (*[]model.TaskInfo, error) {
	var taskInfo []model.TaskInfo
	err := dao.database.Raw(`with st as (
		select coalesce(sum(case when status = 'COMPLETED' then 1 else 0 end), 0) completed,
			   coalesce(sum(case when status = 'COMPLETED' then 0 else 1 end), 0) canceled
		from "transaction"
	)
	select t.id, t.start_date, t.end_date, t.status, t.message, st.completed, st.canceled from st
	left join "transaction" t on true
	order by t.id desc
	limit 5`).Scan(&taskInfo).Error
	if err != nil {
		return nil, err
	}
	return &taskInfo, nil
}
