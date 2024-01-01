package stock

import (
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
	"time"
)

type IStockRepository interface {
	GetSticksWithPage(stockDate time.Time, limit, offset int) (*[]storage.StockFull, error)
}

type Service struct {
	repository IStockRepository
}

func NewStockService(repository IStockRepository) *Service {
	return &Service{
		repository: repository,
	}
}

// Get stocks with the pagginations
func (s *Service) GetStocks(stockDate time.Time, limit, offset int) (*api.StocksFull, error) {
	stocks, err := s.repository.GetSticksWithPage(stockDate, limit, offset)
	if err != nil {
		return nil, err
	}
	size := len(*stocks)
	result := &api.StocksFull{
		Count: 0,
		Items: []api.StockFull{},
	}
	if size == 0 {
		return result, nil
	}
	result.Items = make([]api.StockFull, size)
	for index, item := range *stocks {
		if result.Count == 0 {
			result.Count = item.Total
		}

		result.Items[index] = api.StockFull{
			StockDate:         item.StockDate,
			Source:            item.Source,
			Organization:      item.Org,
			SupplierArticle:   item.SupplierArticle,
			Barcode:           item.Barcode,
			Sku:               item.Sku,
			Name:              item.Name,
			Brand:             item.Brand,
			Warehouse:         item.Warehouse,
			Quantity:          item.Quantity,
			Price:             item.Price,
			PriceWithDiscount: item.PriceWithDiscount,
		}
	}
	return result, nil
}
