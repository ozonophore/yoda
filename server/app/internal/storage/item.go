package storage

import (
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
)

func SaveOrUpdateItem(items *[]model.Item) error {
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

func GetItemCount() int {
	initIfError()
	var count int
	repository.db.Raw(`select count(*) from dl.item`).Scan(&count)
	return count
}
