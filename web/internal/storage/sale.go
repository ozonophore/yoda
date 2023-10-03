package storage

type Sale struct {
	RowNumber         int32   `gorm:"column:rn"`
	ReportYear        int16   `gorm:"column:report_year"`
	ReportMonth       int8    `gorm:"column:report_month"`
	Source            string  `gorm:"column:source"`
	Name              string  `gorm:"column:name"`
	SupplierArticle   string  `gorm:"column:supplier_article"`
	ItemId            string  `gorm:"column:item_id"`
	Barcode           string  `gorm:"column:barcode"`
	ExternalCode      string  `gorm:"column:external_code"`
	Oblast            string  `gorm:"column:oblast"`
	Region            string  `gorm:"column:region"`
	Country           string  `gorm:"column:country"`
	Quantity          int32   `gorm:"column:quantity"`
	TotalPrice        float64 `gorm:"column:total_price"`
	PriceWithDiscount float64 `gorm:"column:price_with_discount"`
	Total             int32   `gorm:"column:total"`
}

func (s *Storage) GetSalesByMonth(year uint16, month uint8) (*[]Sale, error) {
	var sales []Sale
	err := s.db.Raw(`select * from dl.report_sales_by_month
         				 where report_year = ? and report_month = ?
						 order by country, oblast, region, item_id`, year, month).Scan(&sales).Error
	if err != nil {
		return nil, err
	}
	return &sales, nil
}

const sql = `with data as (select
    row_number() over () rn,
    report_year,
    report_month,
    report_date,
    source,
    name,
    supplier_article,
    item_id,
    barcode,
    external_code,
    oblast,
    region,
    country,
    quantity,
    total_price,
    price_with_discount
    from dl.report_sales_by_month
    where report_year = ? and report_month = ?
    )
select rn,
    report_year,
    report_month,
    report_date,
    source,
    name,
    supplier_article,
    item_id,
    barcode,
    external_code,
    oblast,
    region,
    country,
    quantity,
    total_price,
    price_with_discount,
    (select count(1) from data) as total from data`

func (s *Storage) GetSaleByMonthWithPagging(year uint16, month uint8, page int32, size int32) (*[]Sale, error) {
	var sales []Sale
	tx := s.db.Raw(sql+` limit ? offset ?`, year, month, size, (page-1)*size).Scan(&sales)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &sales, nil
}
