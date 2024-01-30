package dictionary

import (
	"context"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
)

type IDictionaryRepository interface {
	GetMarketplaces(ctx context.Context) (*[]storage.Marketplace, error)
	GetClusters(context context.Context, filter *string) (*[]string, error)
	GetPositions(offset int32, limit int32, source []string, filter *string) (*[]storage.Position, error)
	GetWarehousesWithoutPages(source []string, code *string, cluster *string) (*[]storage.Warehouse, error)
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

func (s *DictionaryService) GetDictionary(ctx context.Context) (*api.Dictionaries, error) {
	marketplaces, err := s.repository.GetMarketplaces(ctx)
	if err != nil {
		return nil, err
	}
	mp := make([]api.Marketplace, len(*marketplaces))
	for i, item := range *marketplaces {
		mp[i] = api.Marketplace{
			Code:      item.Code,
			Name:      item.Name,
			ShortName: item.ShortName,
		}
	}
	return &api.Dictionaries{
		Marketplaces: mp,
	}, nil
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
