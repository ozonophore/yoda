package model

import "time"

type OrderPageItem struct {
	ID                int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name              string    `gorm:"column:subject;not null" json:"subject"`
	CreatedDate       time.Time `gorm:"column:transaction_date;not null" json:"transaction_date"`
	OwnerCode         string    `gorm:"column:owner_code;not null" json:"owner_code"`
	Source            string    `gorm:"column:source;not null" json:"source"`
	SupplierArticle   string    `gorm:"column:supplier_article;not null" json:"supplier_article"`
	Barcode           string    `gorm:"column:barcode;not null" json:"barcode"`
	Quantity          int       `gorm:"column:quantity;not null" json:"quantity"`
	WarehouseName     string    `gorm:"column:warehouse_name;not null" json:"warehouse_name"`
	Total             int       `gorm:"column:total;not null" json:"total"`
	TotalPrice        float64   `gorm:"column:total_price" json:"total_price"`                 // Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)
	PriceWithDiscount float64   `gorm:"column:price_with_discount" json:"price_with_discount"` // Цена со скидкой
	Status            string    `gorm:"column:status" json:"status"`
}
