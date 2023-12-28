package storage

import (
	sql2 "database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Position struct {
	Id int32 `gorm:"column:rn"`
	// Code1c Код 1С
	Code1c string `gorm:"column:id"`
	// Name Наименование позиции
	Name    string `gorm:"column:name"`
	Barcode string `gorm:"column:barcode"`
	// Marketplace Торговая точка
	Marketplace string `gorm:"column:marketplace"`
	// MarketplaceId Наименование точки
	MarketplaceId string `gorm:"column:marketplace_id"`
	// Org Организация
	Org   string `gorm:"column:org"`
	Total int32  `gorm:"column:total"`
}

const SQL = `with s as(
				select row_number() over () rn, i.id, i.name, b.barcode, ow.name "org", m.code "marketplace", b.marketplace_id from dl.item i
				inner join dl.barcode b on b.item_id = i.id
				inner  join ml.owner ow on ow.organisation_id = b.organisation_id
				inner join ml.marketplace m on m.marketplace_id = b.marketplace_id
                where m.code in @source %s)
			select rn, id, name, barcode, org, marketplace, marketplace_id, (select count(1) from s) total from s `

func (s *Storage) GetPositions(offset int32, limit int32, source []string, filter *string) (*[]Position, error) {
	var positions []Position

	var filterSQL string
	var filterStr string
	if filter != nil {
		filterStr = "%" + strings.ToUpper(*filter) + "%"
		filterSQL = `and (i.id like @filter or upper(i.name) like @filter or b.barcode like @filter or upper(ow.name) like @filter)`
	} else {
		filterStr = ""
		filterSQL = ""
	}

	var tx *gorm.DB
	tx = s.db.Raw(fmt.Sprintf(SQL, filterSQL)+` limit @limit offset @offset`,
		sql2.Named("limit", limit),
		sql2.Named("offset", offset),
		sql2.Named("source", source),
		sql2.Named("filter", filterStr),
	).Scan(&positions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &positions, nil
}
