package storage

import (
	sql2 "database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	RowNumber       int32   `gorm:"column:rn"`
	Source          string  `gorm:"column:source"`
	SupplierArticle string  `gorm:"column:supplier_article"`
	Code1c          string  `gorm:"column:code_1c"`
	ExternalCode    string  `gorm:"column:external_code"`
	Name            string  `gorm:"column:name"`
	Barcode         string  `gorm:"column:barcode"`
	Brand           string  `gorm:"column:brand"`
	OrderedQuantity int32   `gorm:"column:ordered_quantity"`
	OrderSum        float32 `gorm:"column:order_sum"`
	Balance         int32   `gorm:"column:balance"`
	Total           int32   `gorm:"column:total"`
}

const orderSQL = `with data as (select 
                    row_number() over () rn,
                    source,
					supplier_article,
					code_1c,
					external_code,
					name,
					barcode,
					brand,
					ordered_quantity,
					order_sum,
					balance from dl.report_order_by_day where report_date = @date
                    and %s)
					select 
                    rn,
                    source,
					supplier_article,
					code_1c,
					external_code,
					name,
					barcode,
					brand,
					ordered_quantity,
					order_sum,
					balance, (select count(1) from data) as total from data`

var orderSQLFilter = `and (supplier_article like @filter or code_1c like @filter or external_code like @filter or name like @filter or barcode like @filter or brand like @filter)`

func (s *Storage) GetOrdersByDay(date time.Time) (*[]Order, error) {
	var orders []Order
	sql := fmt.Sprintf(orderSQL, "1=1")
	err := s.db.Raw(sql,
		sql2.Named("date", date.Format(time.DateOnly)),
	).Scan(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (s *Storage) GetOrdersByDayWithPagging(date time.Time, filter string, source string, page int32, size int32) (*[]Order, error) {
	var orders []Order

	where := " 1=1 "
	if source != "" {
		where += orderSQLFilter
	}
	if source != "" {
		where += " and source = @source "
	}

	var tx *gorm.DB
	tx = s.db.Raw(fmt.Sprintf(orderSQL, where)+` limit @size offset @rows`,
		sql2.Named("date", date.Format(time.DateOnly)),
		sql2.Named("size", size),
		sql2.Named("rows", (page-1)*size),
		sql2.Named("filter", "%"+filter+"%"),
		sql2.Named("source", source),
	).Scan(&orders)
	//if len(filter) == 0 {
	//	tx = s.db.Raw(fmt.Sprintf(orderSQL, "")+` limit @size offset @offset`,
	//		sql2.Named("date", date.Format(time.DateOnly)),
	//		sql2.Named("size", size),
	//		sql2.Named("offset", (page-1)*size)).Scan(&orders)
	//} else {
	//	tx = s.db.Raw(orderSQLWithFilter+` limit @size offset @offset`,
	//		sql2.Named("date", date.Format(time.DateOnly)),
	//		sql2.Named("filter", "%"+filter+"%"),
	//		sql2.Named("size", size),
	//		sql2.Named("offset", (page-1)*size)).Scan(&orders)
	//}
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &orders, nil
}
