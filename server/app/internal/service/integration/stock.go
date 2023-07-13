package integration

import (
	"context"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/common/pkg/model"
	"time"
)

type stockRepositoryInterface interface {
	SaveStockFrom1C(stocks *[]model.Stock) error
}

type StockService struct {
	repository stockRepositoryInterface
}

func NewStockService(repository stockRepositoryInterface) *StockService {
	return &StockService{
		repository: repository,
	}
}

func (s *StockService) UploadStocks(ctx context.Context, stocks *[]integration.Stock) error {
	items := mapStocks(stocks)
	return s.repository.SaveStockFrom1C(items)
}

func mapStocks(stocks *[]integration.Stock) *[]model.Stock {
	count := len(*stocks)
	items := make([]model.Stock, count, count)
	stockDate := time.Now().UTC()
	for i, item := range *stocks {
		items[i] = model.Stock{
			StockDate: stockDate,
			ItemID:    item.Id,
			Quantity:  item.Quantity,
		}
	}
	return &items
}
