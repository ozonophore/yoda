package report

import "time"

func (*ReportByItem) TableName() string {
	return "dl.report_by_item"
}

type ReportByItem struct {
	ReportDate      time.Time `gorm:"column:report_date"`
	OwnerCode       string    `gorm:"column:owner_code"`
	Source          string    `gorm:"column:source"`
	SupplierArticle string    `gorm:"column:supplier_article"`
	Barcode         string    `gorm:"column:barcode"`
	ExternalCode    string    `gorm:"column:external_code"`
	Def30           int64     `gorm:"column:def30"`
	DayInStock30    int64     `gorm:"column:day_in_stock30"`
	Def5            int64     `gorm:"column:def5"`
	DayInStock5     int64     `gorm:"column:day_in_stock5"`
	AvgPrice        float64   `gorm:"column:avg_price"`
	MinPrice        float64   `gorm:"column:min_price"`
	MaxPrice        float64   `gorm:"column:max_price"`
	Quantity30      int64     `gorm:"column:quantity30"`
	Quantity5       int64     `gorm:"column:quantity5"`
	OrderByDay30    float64   `gorm:"column:order_by_day30"`
	OrderByDay5     float64   `gorm:"column:order_by_day5"`
	ForecastOrder30 float64   `gorm:"column:forecast_order30"`
	ForecastOrder5  float64   `gorm:"column:forecast_order5"`
	ItemId          string    `gorm:"column:item_id"`
	ItemName        string    `gorm:"column:item_name"`
	MarketplaceId   string    `gorm:"column:marketplace_id"`
	OrgId           string    `gorm:"column:org_id"`
	Quantity        int64     `gorm:"column:quantity"`
	IsExcluded      bool      `gorm:"column:is_excluded"`
	Quantity1С      int64     `gorm:"column:quantity1c"`
}
