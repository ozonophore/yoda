package storage

import (
	sql2 "database/sql"
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
	StockDate         time.Time `gorm:"column:stock_date"`
	Source            string    `gorm:"column:source"`
	Org               string    `gorm:"column:org"`
	SupplierArticle   string    `gorm:"column:supplier_article"`
	Barcode           string    `gorm:"column:barcode"`
	Sku               string    `gorm:"column:sku"`
	Name              string    `gorm:"column:name"`
	Brand             string    `gorm:"column:brand"`
	Warehouse         string    `gorm:"column:warehouse"`
	Quantity          float32   `gorm:"column:quantity"`
	Price             float32   `gorm:"column:price"`
	PriceWithDiscount float32   `gorm:"column:price_with_discount"`
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
where stock_date =@stock_date`

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

func (s *Storage) GetSticksWithPage(stockDate time.Time, limit, offset int) (*[]StockFull, error) {
	var stocks []StockFull
	err := s.db.Raw(sql_stocks,
		sql2.Named("stock_date", stockDate),
		sql2.Named("limit", limit),
		sql2.Named("offset", offset),
	).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}

func (s *Storage) GetStocks(stockDate time.Time) (*[]StockFull, error) {
	var stocks []StockFull
	err := s.db.Raw(sql_stocks_base,
		sql2.Named("stock_date", stockDate),
	).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}
