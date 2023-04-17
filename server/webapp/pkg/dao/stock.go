package dao

import (
	"fmt"
	"github.com/yoda/common/pkg/model"
	"time"
)

func GetPageStocks(limit, offset int, date time.Time, search *string) (*[]model.StockPageItem, error) {
	var stocks []model.StockPageItem
	ifSerchNotNUll := func(search *string) string {
		if search != nil {
			return fmt.Sprintf(`AND ("subject" like '%%%[1]s%%' 
			OR "supplier_article" like '%%%[1]s%%'
            OR "subject" like '%%%[1]s%%'
			OR "barcode" like '%%%[1]s%%'
			OR "warehouse_name" like '%%%[1]s%%'
			OR "owner_code" like '%%%[1]s%%'	
			OR "source" like '%%%[1]s%%'
			)`, *search)
		}
		return ""
	}
	err := dao.database.Raw(fmt.Sprintf(`WITH cte AS (SELECT "id",
                    "transaction_date",
                    "owner_code",
                    "source",
                    "supplier_article",
                    "subject",
                    "barcode",
                    "quantity",
                    "quantity_full",
                    "warehouse_name",
			        "card_created"
             FROM "stock"
             WHERE  date_trunc('day',"transaction_date") = ?
					 %s
             ORDER by 1)
		SELECT *
		FROM (
			 TABLE cte
				 ORDER BY "id"
				 LIMIT ?
				 OFFSET ?) sub
			 LEFT OUTER JOIN (SELECT count(*) FROM cte) c(total) ON true;`,
		ifSerchNotNUll(search),
	), date, limit, offset).Scan(&stocks).Error
	return &stocks, err
}
