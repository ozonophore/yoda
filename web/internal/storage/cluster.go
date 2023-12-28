package storage

import (
	"context"
	"fmt"
	"strings"
)

func (s *Storage) GetClusters(context context.Context, filter *string) (*[]string, error) {
	condition := ""
	if filter != nil {
		condition = "and upper(cluster) like '%" + strings.ToUpper(*filter) + "'"
	}
	sql := fmt.Sprintf(`select distinct cluster from dl.warehouse where 1=1 %s order by cluster`, condition)
	var values []string
	err := s.db.Raw(sql).Scan(&values).Error
	if err != nil {
		return nil, err
	}
	return &values, nil
}
