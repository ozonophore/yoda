package storage

import (
	"github.com/yoda/web/internal/storage/builder"
	"time"
)

type Stock struct {
	OrgId           string `gorm:"column:organisation_id"`
	MarketplaceId   string `gorm:"column:marketplace_id"`
	Barcode         string `gorm:"column:barcode"`
	Quantity        int32  `gorm:"column:quantity"`
	OrgCode         string `gorm:"column:owner_code"`
	MarketplaceCode string `gorm:"column:source"`
}

type StockFull struct {
	StockDate         time.Time `gorm:"column:stock_date" excel:"title:Дата; width:40"`
	Source            string    `gorm:"column:source" excel:"title:МП; width:20"`
	Org               string    `gorm:"column:org" excel:"title:Кабинет; width:40"`
	SupplierArticle   string    `gorm:"column:supplier_article" excel:"title:Артикул поставщика; width:40"`
	Barcode           string    `gorm:"column:barcode" excel:"title:Штрихкод; width:40"`
	Sku               string    `gorm:"column:sku" excel:"title:SKU; width:40"`
	Name              string    `gorm:"column:name" excel:"title:Наименование; width:40"`
	Brand             string    `gorm:"column:brand" excel:"title:Бренд; width:40"`
	Warehouse         string    `gorm:"column:warehouse" excel:"title:Склад; width:40"`
	Quantity          float32   `gorm:"column:quantity" excel:"title:Количество; width:40"`
	Price             float32   `gorm:"column:price" excel:"title:Цена; width:40"`
	PriceWithDiscount float32   `gorm:"column:price_with_discount" excel:"title:Цена со скидкой; width:40"`
	Total             int32     `gorm:"column:total"`
}

func (s *Storage) GetStocksByDate(date time.Time) (*[]Stock, error) {
	stocks := []Stock{}
	err := s.db.Raw(`with t as (select owner_code, source, barcode, round(sum(quantity)) quantity from dl.stock_daily
                                                              where stock_date = ?
                                                              group by owner_code, source, barcode
															  having sum(quantity) > 0)
			select o.name owner_code, o.organisation_id, t.source, m.marketplace_id, barcode, quantity
			from t
			inner join ml.marketplace m on m.code = t.source
			inner join ml.owner o on o.code = t.owner_code`, date.Format(time.DateOnly)).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}

const sql_stocks_base = `select s.stock_date, s.source, o.name org, s.supplier_article, s.barcode, s.external_code sku,
       s.name, s.brand, s.warehouse, s.quantity, s.price, s.price_with_discount
       from bl.stock s
inner join ml.owner o on o.code = s.owner_code
where stock_date =@stock_date and source in @source`

const sql_stocks = `with t as (` + sql_stocks_base + `) 
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

func (s *Storage) GetSticksWithPage(stockDate time.Time, limit, offset int, source *[]string, filter *string) (*[]StockFull, error) {
	var stocks []StockFull

	b := builder.NewStockSQLBuilder(stockDate).Sources(source).Filter(filter).Limit(&limit).Offset(&offset)
	sql, params := b.Build()

	err := s.db.Raw(sql, params).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}

func (s *Storage) GetStocks(stockDate time.Time, source *[]string, filter *string) (*[]StockFull, error) {
	var stocks []StockFull
	b := builder.NewStockSQLBuilder(stockDate).Sources(source).Filter(filter)
	sql, params := b.Build()
	err := s.db.Raw(sql, params).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}
