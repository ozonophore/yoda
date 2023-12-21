package service

import (
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
)

type IDictionaryRepository interface {
	GetPositions(offset int32, limit int32, source []string, filter *string) (*[]storage.Position, error)
	GetWarehouses(offset int32, limit int32, source []string, code *string, cluster *string) (*[]storage.Warehouse, error)
}

type DictionaryService struct {
	repository IDictionaryRepository
}

func NewDictionaryService(repository IDictionaryRepository) *DictionaryService {
	return &DictionaryService{
		repository: repository,
	}
}

func (s *DictionaryService) GetPositions(offset int32, limit int32, source []string, filter *string) (*api.DictPositions, error) {
	positions, err := s.repository.GetPositions(offset, limit, source, filter)
	if err != nil {
		return nil, err
	}
	count := len(*positions)
	if count == 0 {
		return &api.DictPositions{
			Count: 0,
			Items: []api.DictPosition{},
		}, nil
	}
	total := int32(0)
	items := make([]api.DictPosition, count)
	for i, item := range *positions {
		if total == 0 {
			total = item.Total
		}
		items[i] = api.DictPosition{
			Id:            item.Id,
			Code1c:        item.Code1c,
			Name:          item.Name,
			Barcode:       item.Barcode,
			Marketplace:   item.Marketplace,
			MarketplaceId: item.MarketplaceId,
			Org:           item.Org,
		}
	}
	return &api.DictPositions{
		Count: total,
		Items: items,
	}, nil
}

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
