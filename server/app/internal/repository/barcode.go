package repository

import (
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
)

func SaveOrUpdateBarcodes(items *[]model.Barcode) error {
	initIfError()
	return repository.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *items {
			if err := tx.Where(`"barcode_id" =?`, item.BarcodeID).Assign(item).FirstOrCreate(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
