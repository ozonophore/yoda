package model

import (
	"time"
)

const TableNameOrder = `"dl"."order"`

// Order mapped from table <order>
type Order struct {
	ID                int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TransactionDate   time.Time  `gorm:"column:transaction_date;not null" json:"transaction_date"`
	TransactionID     int64      `gorm:"column:transaction_id;not null" json:"transaction_id"`
	Source            string     `gorm:"column:source;not null" json:"source"`
	OwnerCode         string     `gorm:"not null"`
	LastChangeDate    time.Time  `gorm:"column:last_change_date;not null" json:"last_change_date"`                 // Дата обновления информации в сервисе
	LastChangeTime    time.Time  `gorm:"column:last_change_time;not null;serializer:time" json:"last_change_time"` // Время обновления информации в сервисе
	OrderDate         time.Time  `gorm:"column:order_date;not null" json:"order_date"`                             // Дата заказа
	OrderTime         time.Time  `gorm:"column:order_time;not null;serializer:time" json:"order_time"`             // Время заказа
	SupplierArticle   *string    `gorm:"column:supplier_article" json:"supplier_article"`                          // Артикул поставщика
	TechSize          *string    `gorm:"column:tech_size" json:"tech_size"`                                        // Размер
	Barcode           *string    `gorm:"column:barcode" json:"barcode"`                                            // Бар-код
	TotalPrice        float64    `gorm:"column:total_price" json:"total_price"`                                    // Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)
	DiscountPercent   float64    `gorm:"column:discount_percent" json:"discount_percent"`                          // Согласованный итоговый дисконт(процент)
	DiscountValue     float64    `gorm:"column:discount_value" json:"discount_value"`                              // Согласованный итоговый дисконт(значение)
	PriceWithDiscount float64    `gorm:"column:price_with_discount" json:"price_with_discount"`                    // Цена со скидкой
	WarehouseName     *string    `gorm:"column:warehouse_name" json:"warehouse_name"`                              // Название склада отгрузки
	Oblast            *string    `gorm:"column:oblast" json:"oblast"`                                              // Область
	IncomeID          *int64     `gorm:"column:income_id" json:"income_id"`                                        // Номер поставки (от продавца на склад)
	ExternalCode      string     `gorm:"column:external_code" json:"external_code"`                                // Код WB
	Odid              *int64     `gorm:"column:odid" json:"odid"`                                                  // Уникальный идентификатор позиции заказа
	Subject           *string    `gorm:"column:subject" json:"subject"`                                            // Предмет
	Category          *string    `gorm:"column:category" json:"category"`                                          // Категория
	Brand             *string    `gorm:"column:brand" json:"brand"`                                                // Бренд
	IsCancel          bool       `gorm:"column:is_cancel" json:"is_cancel"`                                        // Отмена заказа. true - заказ отменен до оплаты
	Status            string     `gorm:"column:status" json:"status"`                                              // Статус заказа
	CancelDt          *time.Time `gorm:"column:cancel_dt" json:"cancel_dt"`                                        // Дата и время отмены заказа
	GNumber           *string    `gorm:"column:g_number" json:"g_number"`                                          // Номер заказа. Объединяет все позиции одного заказа.
	Sticker           *string    `gorm:"column:sticker" json:"sticker"`                                            // Цифровое значение стикера
	Srid              *string    `gorm:"column:srid" json:"srid"`                                                  // Уникальный идентификатор заказа
	Quantity          int64      `gorm:"column:quantity" json:"quantity"`                                          // Количество
	ItemId            *string    `gorm:"column:item_id" json:"item_id"`                                            // Уникальный идентификатор товара
	BarcodeId         *string    `gorm:"column:barcode_id" json:"barcode_id"`                                      // Уникальный идентификатор штрихкода
	Message           *string    `gorm:"column:message" json:"message"`
	Country           *string    `gorm:"column:country" json:"country"` // Старана доставки
	Region            *string    `gorm:"column:region" json:"region"`   // Регион доставки
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
