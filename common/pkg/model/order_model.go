package model

import (
	"time"
)

const TableNameOrder = "order"

// Order mapped from table <order>
type Order struct {
	ID              int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TransactionID   int64      `gorm:"column:transaction_id;not null" json:"transaction_id"`
	Source          string     `gorm:"column:source;not null" json:"source"`
	LastChangeDate  time.Time  `gorm:"column:last_change_date;not null" json:"last_change_date"`                 // Дата обновления информации в сервисе
	LastChangeTime  time.Time  `gorm:"column:last_change_time;not null;serializer:time" json:"last_change_time"` // Время обновления информации в сервисе
	OrderDate       time.Time  `gorm:"column:order_date;not null" json:"order_date"`                             // Дата заказа
	OrderTime       time.Time  `gorm:"column:order_time;not null;serializer:time" json:"order_time"`             // Время заказа
	SupplierArticle *string    `gorm:"column:supplier_article" json:"supplier_article"`                          // Артикул поставщика
	TechSize        *string    `gorm:"column:tech_size" json:"tech_size"`                                        // Размер
	Barcode         *string    `gorm:"column:barcode" json:"barcode"`                                            // Бар-код
	TotalPrice      float64    `gorm:"column:total_price" json:"total_price"`                                    // Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)
	DiscountPercent *int       `gorm:"column:discount_percent" json:"discount_percent"`                          // Согласованный итоговый дисконт
	WarehouseName   *string    `gorm:"column:warehouse_name" json:"warehouse_name"`                              // Название склада отгрузки
	Oblast          *string    `gorm:"column:oblast" json:"oblast"`                                              // Область
	IncomeID        *int       `gorm:"column:income_id" json:"income_id"`                                        // Номер поставки (от продавца на склад)
	ExternalCode    string     `gorm:"column:external_code" json:"external_code"`                                // Код WB
	Odid            *int32     `gorm:"column:odid" json:"odid"`                                                  // Уникальный идентификатор позиции заказа
	Subject         *string    `gorm:"column:subject" json:"subject"`                                            // Предмет
	Category        *string    `gorm:"column:category" json:"category"`                                          // Категория
	Brand           *string    `gorm:"column:brand" json:"brand"`                                                // Бренд
	IsCancel        bool       `gorm:"column:is_cancel" json:"is_cancel"`                                        // Отмена заказа. true - заказ отменен до оплаты
	CancelDt        *time.Time `gorm:"column:cancel_dt" json:"cancel_dt"`                                        // Дата и время отмены заказа
	GNumber         *string    `gorm:"column:g_number" json:"g_number"`                                          // Номер заказа. Объединяет все позиции одного заказа.
	Sticker         *string    `gorm:"column:sticker" json:"sticker"`                                            // Цифровое значение стикера
	Srid            *string    `gorm:"column:srid" json:"srid"`                                                  // Уникальный идентификатор заказа
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
