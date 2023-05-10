package repository

import (
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
)

func SaveOrUpdateMarketplaces(items *[]model.Marketplace) error {
	initIfError()
	return repository.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *items {
			if err := tx.Where(`"id" =?`, item.ID).Assign(item).FirstOrCreate(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
