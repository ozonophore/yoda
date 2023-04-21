package dao

import "github.com/yoda/webapp/pkg/model"

func GetWeekSales() (*[]model.WeekSales, error) {
	var weekSales []model.WeekSales
	err := dao.database.Raw(`select order_date, price from (select order_date, sum(price_with_discount) * sum(quantity) price
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
