// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameJobMarketplace = "job_marketplace"

// JobMarketplace mapped from table <job_marketplace>
type JobParameter struct {
	JobID    int32   `gorm:"column:job_id;not null" json:"job_id"`
	Source   string  `gorm:"column:source;not null" json:"source"`
	Host     string  `gorm:"column:host;not null" json:"host"`
	Password *string `gorm:"column:password" json:"password"`
	ClientID *string `gorm:"column:client_id" json:"client_id"`
}

// TableName JobMarketplace's table name
func (*JobParameter) TableName() string {
	return TableNameJobMarketplace
}
