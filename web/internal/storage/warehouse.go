package storage

import (
	sql2 "database/sql"
	"fmt"
	"github.com/yoda/web/internal/api"
	"strings"
)

type Warehouse struct {
	Code    string `gorm:"column:code"`
	Cluster string `gorm:"column:cluster"`
	Source  string `gorm:"column:source"`
	Total   int32  `gorm:"column:total"`
}

const WH_SQL = `select code, cluster, source from dl.warehouse 
                             where 1 = 1 %s`

const WH_SQL_WITH_PAGE = `with wh as (
    ` + WH_SQL + `
)
select wh.code, wh.source, wh.cluster, (select count(1) from wh) as total from wh
limit @limit offset @offset;`

func (s *Storage) GetWarehouses(offset int32, limit int32, source []string, code *string, cluster *string) (*[]Warehouse, error) {
	var warehouses []Warehouse
	filter := ""
	if code != nil && len(*code) > 0 {
		filter += fmt.Sprintf(" and upper(code) like '%s'", "%"+strings.ToUpper(*code)+"%")
	}
	if cluster != nil && len(*cluster) > 0 {
		filter += fmt.Sprintf(" and upper(cluster) like '%s'", "%"+strings.ToUpper(*cluster)+"%")
	}
	if len(source) > 0 {
		filter += " and source in @source"
	}

	tx := s.db.Raw(fmt.Sprintf(WH_SQL_WITH_PAGE, filter),
		sql2.Named("limit", limit),
		sql2.Named("offset", offset),
		sql2.Named("source", source),
	).Scan(&warehouses)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &warehouses, nil
}

func (s *Storage) GetWarehousesWithoutPages(source []string, code *string, cluster *string) (*[]Warehouse, error) {
	var warehouses []Warehouse
	filter := ""
	if code != nil && len(*code) > 0 {
		filter += fmt.Sprintf(" and upper(code) like '%s'", "%"+strings.ToUpper(*code)+"%")
	}
	if cluster != nil && len(*cluster) > 0 {
		filter += fmt.Sprintf(" and upper(cluster) like '%s'", "%"+strings.ToUpper(*cluster)+"%")
	}
	if len(source) > 0 {
		filter += " and source in @source"
	}

	tx := s.db.Raw(fmt.Sprintf(WH_SQL, filter),
		sql2.Named("source", source),
	).Scan(&warehouses)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &warehouses, nil
}

func (s *Storage) UpdateWarehouse(value *api.Warehouse) (*api.Warehouse, error) {
	err := s.db.Table("dl.warehouse").
		Where("code=? and source=?", value.Code, value.Source).
		Update("cluster", value.Cluster).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}
