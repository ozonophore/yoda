package storage

import (
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm/clause"
)

func SaveOrUpdateBarcodes(items *[]model.Barcode) error {
	initIfError()
	return repository.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "barcode_id"}},
		DoNothing: true,
	}).CreateInBatches(items, len(*items)).Error
}

func GetBarcodeCount() int {
	initIfError()
	var count int
	repository.db.Raw(`select count(*) from dl.barcode`).Scan(&count)
	return count
}

func GetBarcodeDictionary() (*[]model.OrgSrcBarcode, error) {
	initIfError()
	var item []model.OrgSrcBarcode
	err := repository.db.Raw(`select ow.code as org_code, m.code as source, b.barcode, barcode_id, i.id as item_id, i.name from dl.item i
				inner join dl.barcode b on i.id = b.item_id
				inner join dl.organisation o on b.organisation_id = o.id
				inner join ml.owner ow on ow.organisation_id = o.id
				inner join ml.marketplace m on m.marketplace_id = b.marketplace_id`).Find(&item).Error
	return &item, err
}
