package model

import "time"

const TableNameItem = `"dl"."item"`

// Item mapped from table <item>
type Item struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`             // ID
	Name      string    `gorm:"column:name;not null" json:"name"`           // Name
	UpdatedAt time.Time `gorm:"column:update_at;not null" json:"update_at"` // Updated at
}

// TableName Item's table name
func (*Item) TableName() string {
	return TableNameItem
}
