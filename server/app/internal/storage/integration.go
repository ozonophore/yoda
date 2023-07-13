package storage

import (
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
)

func (r *Repository) SaveStockFrom1C(stocks *[]model.Stock) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(stocks, r.config.BatchSize).Error
	})
}
