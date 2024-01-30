package stock

import (
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/service"
	"github.com/yoda/web/internal/storage"
	"net/http"
	"time"
)

var header []service.ExcelHeaderColumn = []service.ExcelHeaderColumn{
	{
		Title: "МП",
		Width: 20,
		Field: "Source",
	}, {
		Title: "Дата",
		Width: 40,
		Field: "StockDate",
	}, {
		Title: "Кабинет",
		Width: 40,
		Field: "Org",
	}, {
		Title: "Артикул поставщика",
		Width: 40,
		Field: "SupplierArticle",
	}, {
		Title: "Штрихкод",
		Width: 40,
		Field: "Barcode",
	}, {
		Title: "SKU",
		Width: 40,
		Field: "Sku",
	}, {
		Title: "Наименование",
		Width: 40,
		Field: "Name",
	}, {
		Title: "Бренд",
		Width: 40,
		Field: "Brand",
	}, {
		Title: "Склад",
		Width: 40,
		Field: "Warehouse",
	}, {
		Title: "Количество",
		Width: 40,
		Field: "Quantity",
	}, {
		Title: "Цена",
		Width: 40,
		Field: "Price",
	}, {
		Title: "Цена со скидкой",
		Width: 40,
		Field: "PriceWithDiscount",
	},
}

type IStockRepository interface {
	GetSticksWithPage(stockDate time.Time, limit, offset int, source *[]string, filter *string) (*[]storage.StockFull, error)
	GetStocks(stockDate time.Time, source *[]string, filter *string) (*[]storage.StockFull, error)
}

type Service struct {
	repository IStockRepository
}

func NewStockService(repository IStockRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ExportStocks(writer http.ResponseWriter, stockDate time.Time, source *[]string, filter *string) error {
	stocks, err := s.repository.GetStocks(stockDate, source, filter)
	if err != nil {
		return err
	}
	return service.GenerateExcelDoc(writer, "Заказы", stocks, &header)
}

// Get stocks with the pagginations
func (s *Service) GetStocks(stockDate time.Time, limit, offset int, source *[]string, filter *string) (*api.StocksFull, error) {
	stocks, err := s.repository.GetSticksWithPage(stockDate, limit, offset, source, filter)
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
