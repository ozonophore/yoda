package service

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/xuri/excelize/v2"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/storage"
	"io"
	"time"
)

type column struct {
	title string
	width float64
	field string
}

var headers = []column{
	{
		title: "МП",
		width: 20,
		field: "Source",
	},
	{
		title: "Бренд",
		width: 20,
		field: "Brand",
	},
	{
		title: "Наименование",
		width: 30,
		field: "Name",
	},
	{
		title: "Артикул продавца",
		width: 20,
		field: "SupplierArticle",
	},
	{
		title: "Артикул МП",
		width: 20,
		field: "ExternalCode",
	},
	{
		title: "Код 1С",
		width: 20,
		field: "Code1c",
	},
	{
		title: "Баркод",
		width: 20,
		field: "Barcode",
	},
	{
		title: "Заказано шт",
		width: 20,
		field: "OrderedQuantity",
	},
	{
		title: "Сумма заказов",
		width: 20,
		field: "OrderSum",
	},
	{
		title: "Текущий остаток, шт",
		width: 20,
		field: "Balance",
	},
}

type OrderRepositoryInterface interface {
	GetOrdersByDay(date time.Time) (*[]storage.Order, error)
	GetOrdersByDayWithPagging(date time.Time, filter string, source string, page int32, size int32) (*[]storage.Order, error)
}

type OrderService struct {
	repository OrderRepositoryInterface
}

func NewOrderService(repository OrderRepositoryInterface) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) PrepareAndReturnExcel(writer io.Writer, date time.Time) error {
	f := excelize.NewFile()
	defer f.Close()

	headerStyle, _ := f.NewStyle(GetHeaderStyle())
	sheetIndex := f.GetActiveSheetIndex()
	sheetName := fmt.Sprintf("Заказы за %s", date.Format(time.DateOnly))
	f.SetSheetName(f.GetSheetName(sheetIndex), sheetName)

	for index, header := range headers {
		colChar := ToChar(index)
		cell := fmt.Sprintf("%s3", colChar)
		f.SetCellValue(sheetName, cell, header.title)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
		f.SetColWidth(sheetName, "A", colChar, header.width)
	}

	orders, err := s.repository.GetOrdersByDay(date)
	if err != nil {
		return err
	}
	for rowIndex, order := range *orders {
		row := rowIndex + 4
		fields := structs.Map(&order)
		for index, header := range headers {
			colChar := ToChar(index)
			cell := fmt.Sprintf("%s%d", colChar, row)
			field := fields[header.field]
			f.SetCellValue(sheetName, cell, field)
		}
	}

	lastColName := ToChar(len(headers) - 1)
	f.AutoFilter(sheetName, fmt.Sprintf("A3:%s3", lastColName), []excelize.AutoFilterOptions{})

	return f.Write(writer)
}

func (s *OrderService) GetOrders(date time.Time, filter string, source string, page int32, size int32) (*api.Orders, error) {
	orders, err := s.repository.GetOrdersByDayWithPagging(date, filter, source, page, size)
	if err != nil {
		return nil, err
	}
	len := len(*orders)
	if len == 0 {
		return &api.Orders{
			Items: []api.Order{},
			Count: 0,
		}, nil
	}
	total := int32(0)
	items := make([]api.Order, len)
	for i, order := range *orders {
		if total == 0 {
			total = order.Total
		}
		items[i] = api.Order{
			Id:              order.RowNumber,
			Source:          order.Source,
			Brand:           order.Brand,
			Name:            order.Name,
			SupplierArticle: order.SupplierArticle,
			Barcode:         order.Barcode,
			ExternalCode:    order.ExternalCode,
			Code1c:          order.Code1c,
			OrderedQuantity: order.OrderedQuantity,
			OrderSum:        order.OrderSum,
			Balance:         order.Balance,
		}
	}
	return &api.Orders{
		Items: items,
		Count: total,
	}, nil
}
