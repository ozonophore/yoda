package storage

import (
	sql2 "database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
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
	OwnerCode       string  `gorm:"column:owner_code"`
	OrgName         string  `gorm:"column:org_name"`
}

type OrderProduct struct {
	RowNumber         int32     `gorm:"column:row_num"`
	OrderDate         time.Time `gorm:"column:order_date"`
	Source            string    `gorm:"column:source"`
	SupplierArticle   string    `gorm:"column:supplier_article"`
	Barcode           string    `gorm:"column:barcode"`
	Brand             string    `gorm:"column:brand"`
	OrgName           string    `gorm:"column:org_name"`
	ExternalCode      string    `gorm:"column:external_code"`
	Total             int32     `gorm:"column:total"`
	ItemID            string    `gorm:"column:item_id"`
	ItemName          string    `gorm:"column:item_name"`
	Quantity          int32     `gorm:"column:quantity"`
	QuantityCanceled  int32     `gorm:"column:quantity_canceled"`
	QuantityDelivered int32     `gorm:"column:quantity_delivered"`
}

const orderProductSQL = `select
    row_number() over () row_num,
    op.order_date,
    op.source,
    ow.name org_name,
    op.supplier_article,
    op.barcode,
    op.brand,
    op.external_code,
    i.id item_id,
    i.name item_name,
    op.quantity,
    op.quantity_canceled,
    op.quantity_delivered
from bl.order_position op
inner join ml.marketplace m on m.code = op.source
inner join ml.owner ow on ow.code = op.owner_code
left outer join dl.barcode b on b.barcode = op.barcode and b.marketplace_id = m.marketplace_id and b.organisation_id= ow.organisation_id
left outer join dl.item i on i.id = b.item_id
where op.order_date >= @dateFrom and op.order_date <= @dateTo
and %s
order by op.order_date`

const SQL_PRODUCT_WITHOUT_PAGE = `with data as (select
    row_number() over () row_num,
    op.order_date,
    op.source,
    ow.name org_name,
    op.supplier_article,
    op.barcode,
    op.brand,
    op.external_code,
    i.id item_id,
    i.name item_name,
    op.quantity,
    op.quantity_canceled,
    op.quantity_delivered
from bl.order_position op
inner join ml.marketplace m on m.code = op.source
inner join ml.owner ow on ow.code = op.owner_code
left outer join dl.barcode b on b.barcode = op.barcode and b.marketplace_id = m.marketplace_id and b.organisation_id= ow.organisation_id
left outer join dl.item i on i.id = b.item_id
where op.order_date >= @dateFrom and op.order_date <= @dateTo
and %s
order by op.order_date) 
select row_num, order_date, source, org_name, supplier_article, barcode, brand
,external_code, (select count(1) from data) as total, item_id, item_name, quantity
,quantity_canceled, quantity_delivered
from data`

const SQL_PRODUCT = SQL_PRODUCT_WITHOUT_PAGE + ` limit @limit offset @offset`

const SQl_WITH_GROUP_WITHOUT_PAGE = `with data as (select
    row_number() over () row_num,
    op.source,
    ow.name org_name,
    op.supplier_article,
    op.barcode,
    op.brand,
    op.external_code,
    i.id item_id,
    i.name item_name,
    sum(op.quantity) quantity,
    sum(op.quantity_canceled) quantity_canceled,
    sum(op.quantity_delivered) quantity_delivered
from bl.order_position op
inner join ml.marketplace m on m.code = op.source
inner join ml.owner ow on ow.code = op.owner_code
left outer join dl.barcode b on b.barcode = op.barcode and b.marketplace_id = m.marketplace_id and b.organisation_id= ow.organisation_id
left outer join dl.item i on i.id = b.item_id
where op.order_date >= @dateFrom and op.order_date <= @dateTo
and %s
group by op.source, ow.name, op.supplier_article, op.barcode, op.brand, op.external_code, i.id, i.name)
select row_num, source, org_name, supplier_article, barcode, brand
,external_code, (select count(1) from data) as total, item_id, item_name, quantity
,quantity_canceled, quantity_delivered
from data`

const SQl_WITH_GROUP = SQl_WITH_GROUP_WITHOUT_PAGE +
	` limit @limit offset @offset`

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
                    org_name,
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
				    org_name,
					balance, (select count(1) from data) as total from data`

var orderSQLFilter = `and (supplier_article like @filter or 
                      code_1c like @filter or external_code like @filter 
                      or org_name like @filter
                      or name like @filter or barcode like @filter or brand like @filter)`

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
	if filter != "" {
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

func (s *Storage) GetOrdersProductWithoutPage(dateFrom time.Time, dateTo time.Time, filter *string, groupBy *string) (*[]OrderProduct, error) {
	var orders []OrderProduct
	var filterSQL *string
	if filter != nil {
		value := *filter
		s := fmt.Sprintf(` (upper(op.source) like '%[1]v' or upper(op.supplier_article) like '%[1]v' or upper(op.barcode) like '%[1]v' 
            or upper(op.brand) like '%[1]v' or upper(ow.name) like '%[1]v' `+
			`or upper(op.external_code::text) like '%[1]v' or upper(i.name) like '%[1]v' or upper(i.id) like '%[1]v')`, "%"+strings.ToUpper(value)+"%")
		filterSQL = &s
	} else {
		s := `1 = 1`
		filterSQL = &s
	}
	var tx *gorm.DB
	if groupBy == nil {
		tx = s.db.Raw(fmt.Sprintf(SQL_PRODUCT_WITHOUT_PAGE, *filterSQL),
			sql2.Named("dateFrom", dateFrom.Format(time.DateOnly)),
			sql2.Named("dateTo", dateTo.Format(time.DateOnly)),
		).Scan(&orders)
	} else {
		tx = s.db.Raw(fmt.Sprintf(SQl_WITH_GROUP_WITHOUT_PAGE, *filterSQL),
			sql2.Named("dateFrom", dateFrom.Format(time.DateOnly)),
			sql2.Named("dateTo", dateTo.Format(time.DateOnly)),
		).Scan(&orders)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &orders, nil
}

func (s *Storage) GetOrdersProduct(dateFrom time.Time, dateTo time.Time, filter *string, offset int32, limit int32, groupBy *string) (*[]OrderProduct, error) {
	var orders []OrderProduct
	var filterSQL *string
	if filter != nil {
		v := *filter
		s := fmt.Sprintf(` (upper(op.source) like '%[1]v' or upper(op.supplier_article) like '%[1]v' or upper(op.barcode) like '%[1]v' 
            or upper(op.brand) like '%[1]v' or upper(ow.name) like '%[1]v' `+
			`or upper(op.external_code::text) like '%[1]v' or upper(i.name) like '%[1]v' or upper(i.id) like '%[1]v')`, "%"+strings.ToUpper(v)+"%")
		filterSQL = &s
	} else {
		s := `1 = 1`
		filterSQL = &s
	}
	var tx *gorm.DB
	if groupBy == nil {
		tx = s.db.Raw(fmt.Sprintf(SQL_PRODUCT, *filterSQL),
			sql2.Named("dateFrom", dateFrom.Format(time.DateOnly)),
			sql2.Named("dateTo", dateTo.Format(time.DateOnly)),
			sql2.Named("offset", offset),
			sql2.Named("limit", limit),
		).Scan(&orders)
	} else {
		tx = s.db.Raw(fmt.Sprintf(SQl_WITH_GROUP, *filterSQL),
			sql2.Named("dateFrom", dateFrom.Format(time.DateOnly)),
			sql2.Named("dateTo", dateTo.Format(time.DateOnly)),
			sql2.Named("offset", offset),
			sql2.Named("limit", limit),
		).Scan(&orders)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &orders, nil
}
