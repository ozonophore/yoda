package builder

import (
	"fmt"
	"time"
)

const sql_stocks_base = `select s.stock_date, s.source, o.name org, s.supplier_article, s.barcode, s.external_code sku,
       s.name, s.brand, s.warehouse, s.quantity, s.price, s.price_with_discount
       from bl.stock s
inner join ml.owner o on o.code = s.owner_code
where stock_date = @stock_date`
const sql_stock_where_source = ` and source in @source`
const sql_stock_where_filter = ` and (s.supplier_article like @filter or s.barcode like @filter or s.external_code like @filter or s.name like @filter or s.brand like @filter or s.warehouse like @filter)`

const sql_wrapper = `with t as (%s) 
select 
stock_date, 
source, 
org, 
supplier_article, 
barcode, 
sku, 
name, 
brand, 
warehouse, 
quantity, 
price, 
price_with_discount, 
(select count(1) from t) total
from t
limit @limit offset @offset`

type StockSQLBuilder struct {
	sources   *[]string
	filter    *string
	stockSate time.Time
	limit     *int
	offset    *int
}

func NewStockSQLBuilder(stockDate time.Time) *StockSQLBuilder {
	return &StockSQLBuilder{
		stockSate: stockDate,
	}
}

func (s *StockSQLBuilder) Sources(sources *[]string) *StockSQLBuilder {
	s.sources = sources
	return s
}

func (s *StockSQLBuilder) Filter(filter *string) *StockSQLBuilder {
	s.filter = filter
	return s
}

func (s *StockSQLBuilder) Limit(limit *int) *StockSQLBuilder {
	s.limit = limit
	return s
}

func (s *StockSQLBuilder) Offset(offset *int) *StockSQLBuilder {
	s.offset = offset
	return s
}

func (s *StockSQLBuilder) Build() (string, map[string]interface{}) {
	var vars = make(map[string]interface{})
	vars["stock_date"] = s.stockSate
	query := sql_stocks_base
	if s.sources != nil && len(*s.sources) != 0 {
		vars["source"] = *s.sources
	} else {
		vars["source"] = []string{}
	}
	query += sql_stock_where_source
	if s.filter != nil && len(*s.filter) != 0 {
		str := *s.filter
		vars["filter"] = "%" + str + "%"
		query += sql_stock_where_filter
	}
	if s.limit != nil || s.offset != nil {
		query = fmt.Sprintf(sql_wrapper, query)
		vars["limit"] = *s.limit
		vars["offset"] = *s.offset
	}

	return query, vars
}
