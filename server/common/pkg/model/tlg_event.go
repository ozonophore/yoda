// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTlgEvent = "tlg_event"

// TlgEvent mapped from table <tlg_event>
type TlgEvent struct {
	ChatID   int64  `gorm:"column:chat_id;primaryKey" json:"chat_id"`
	DataType string `gorm:"column:data_type;primaryKey" json:"data_type"`
}

// TableName TlgEvent's table name
func (*TlgEvent) TableName() string {
	return TableNameTlgEvent
}
