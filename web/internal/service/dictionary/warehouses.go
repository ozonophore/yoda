package dictionary

import (
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/service"
	"net/http"
)

func (s *DictionaryService) GetWarehouses(offset int32, limit int32, source []string, code *string, cluster *string) (*api.Warehouses, error) {
	wh, err := s.repository.GetWarehouses(offset, limit, source, code, cluster)
	if err != nil {
		return nil, err
	}
	count := int32(len(*wh))
	if count == 0 {
		return &api.Warehouses{
			Count: 0,
			Items: []api.Warehouse{},
		}, nil
	}
	items := make([]api.Warehouse, count)

	for i, item := range *wh {
		count = item.Total
		items[i] = api.Warehouse{
			Code:    item.Code,
			Cluster: item.Cluster,
			Source:  item.Source,
		}
	}
	return &api.Warehouses{
		Count: count,
		Items: items,
	}, nil
}

var headers = []service.ExcelHeaderColumn{
	{
		Title: "Кластер",
		Width: 120,
		Field: "Cluster",
	}, {
		Title: "Код склада",
		Width: 140,
		Field: "Code",
	}, {
		Title: "МП",
		Width: 70,
		Field: "Source",
	},
}

func (s *DictionaryService) ExportWarehouses(writer http.ResponseWriter, source []string, code *string, cluster *string) error {
	wh, err := s.repository.GetWarehousesWithoutPages(source, code, cluster)
	if err != nil {
		return err
	}
	return service.GenerateExcelDoc(writer, "Заказы", wh, &headers)
}
