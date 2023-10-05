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
		title: "Юр. лицо",
		width: 25,
		field: "OrgName",
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

type sheetType struct {
	Row  int
	Name string
}

func (s *OrderService) PrepareAndReturnExcel(writer io.Writer, date time.Time) error {
	f := excelize.NewFile()
	defer f.Close()

	headerStyle, _ := f.NewStyle(GetHeaderStyle())
	sheetIndex := f.GetActiveSheetIndex()

	sheets := make(map[string]*sheetType, 2)

	orders, err := s.repository.GetOrdersByDay(date)
	if err != nil {
		return err
	}
	for _, order := range *orders {
		sheet, ok := sheets[order.Source]
		if !ok {
			sheetName := fmt.Sprintf("Заказы %s за %s", order.Source, date.Format(time.DateOnly))
			sheet = &sheetType{
				Row:  4,
				Name: sheetName,
			}
			sheets[order.Source] = sheet
			if len(sheets) == 1 {
				f.SetSheetName(f.GetSheetName(sheetIndex), sheetName)
			} else {
				f.NewSheet(sheetName)
			}
			s.createHeaders(f, sheetName, headerStyle)
		}
		fields := structs.Map(&order)
		for index, header := range headers {
			colChar := ToChar(index)
			cell := fmt.Sprintf("%s%d", colChar, sheet.Row)
			field := fields[header.field]
			f.SetCellValue(sheet.Name, cell, field)
		}
		sheet.Row++
	}

	lastColName := ToChar(len(headers) - 1)
	for _, sheet := range sheets {
		f.AutoFilter(sheet.Name, fmt.Sprintf("A3:%s3", lastColName), []excelize.AutoFilterOptions{})
	}

	return f.Write(writer)
}

func (s *OrderService) createHeaders(f *excelize.File, sheetName string, headerStyle int) {
	for index, header := range headers {
		colChar := ToChar(index)
		cell := fmt.Sprintf("%s3", colChar)
		f.SetCellValue(sheetName, cell, header.title)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
		f.SetColWidth(sheetName, "A", colChar, header.width)
	}
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
			Org:             order.OrgName,
		}
	}
	return &api.Orders{
		Items: items,
		Count: total,
	}, nil
}
