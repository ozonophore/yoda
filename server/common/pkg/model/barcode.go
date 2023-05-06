// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBarcode = `"dl"."barcode"`

// Barcode mapped from table <barcode>
type Barcode struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`                         // ID
	Barcode        string    `gorm:"column:barcode;not null" json:"barcode"`                 // Barcode
	OrganisationID string    `gorm:"column:organisation_id;not null" json:"organisation_id"` // Organization ID
	MarketplaceID  string    `gorm:"column:marketplace_id;not null" json:"marketplace_id"`   // Marketplace ID
	Article        *string   `gorm:"column:article" json:"article"`                          // Article
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`           // Updated at
}

// TableName Barcode's table name
func (*Barcode) TableName() string {
	return TableNameBarcode
}
