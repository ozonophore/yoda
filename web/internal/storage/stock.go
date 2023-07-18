package storage

import "time"

type Stock struct {
	OrgId           string `gorm:"column:organisation_id"`
	MarketplaceId   string `gorm:"column:marketplace_id"`
	Barcode         string `gorm:"column:barcode"`
	Quantity        int32  `gorm:"column:quantity"`
	OrgCode         string `gorm:"column:owner_code"`
	MarketplaceCode string `gorm:"column:source"`
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
			inner join ml.owner o on o.code = t.owner_code
			inner join ml.owner o on o.code = t.owner_code`, date.Format(time.DateOnly)).Scan(&stocks).Error
	if err != nil {
		return nil, err
	}
	return &stocks, nil
}
