package sale

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/xuri/excelize/v2"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/service"
	"github.com/yoda/web/internal/storage"
	"io"
)

var monthes = []string{
	"январь",
	"февраль",
	"март",
	"апрель",
	"май",
	"июнь",
	"июль",
	"август",
	"сентябрь",
	"октябрь",
	"ноябрь",
	"декабрь",
}

type regionKey struct {
	Oblast  string
	Region  string
	Country string
}

type productKey struct {
	source          string
	name            string
	supplierArticle string
	itemId          string
	barcode         string
	externalCode    string
}

type column struct {
	group    *string
	subgroup *string
	title    string
	width    float64
	field    string
}

var headers = []column{
	{
		title: "МП",
		width: 10,
		field: "Source",
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
		field: "ItemId",
	},
	{
		title: "Баркод",
		width: 20,
		field: "Barcode",
	},
}

var measures = []string{
	"Кол-во",
	"Сумма",
	"Сумма со скидкой",
}

type SaleRepositoryInterface interface {
	GetSalesByMonth(year uint16, month uint8) (*[]storage.Sale, error)
	GetSaleByMonthWithPagging(year uint16, month uint8, page int32, size int32) (*[]storage.Sale, error)
}

type SaleService struct {
	repository SaleRepositoryInterface
}

func NewSaleService(repository SaleRepositoryInterface) *SaleService {
	return &SaleService{
		repository: repository,
	}
}

func (s *SaleService) PrepareAndReturnExcel(writer io.Writer, year uint16, month uint8) error {
	f := excelize.NewFile()
	defer f.Close()

	headerStyle, _ := f.NewStyle(service.GetHeaderStyle())
	sheetIndex := f.GetActiveSheetIndex()
	sheetName := fmt.Sprintf("Продажи за %s %d", monthes[month-1], year)
	f.SetSheetName(f.GetSheetName(sheetIndex), sheetName)

	for index, header := range headers {
		colChar := service.ToChar(index)
		cell1 := fmt.Sprintf("%s2", colChar)
		cell2 := fmt.Sprintf("%s5", colChar)
		f.MergeCell(sheetName, cell1, cell2)
		f.SetCellValue(sheetName, cell1, header.title)
		f.SetCellStyle(sheetName, cell1, cell2, headerStyle)
		f.SetColWidth(sheetName, "A", colChar, header.width)
	}

	sales, err := s.repository.GetSalesByMonth(year, month)
	if err != nil {
		return err
	}

	index := 0
	customHeader := make(map[regionKey]*uint16)
	prevCountry := ""
	prevOblast := ""
	prevRegion := ""
	productRow := make(map[productKey]uint32)
	row := 0
	for _, sale := range *sales {
		key := regionKey{
			Oblast:  sale.Oblast,
			Region:  sale.Region,
			Country: sale.Country,
		}
		sales := customHeader[key]
		if sales == nil {
			i := uint16(index)
			customHeader[key] = &i

			for _, measure := range measures {

				colChar := service.ToChar(index + len(headers))
				cell1 := fmt.Sprintf("%s2", colChar)
				cell2 := fmt.Sprintf("%s3", colChar)
				cell3 := fmt.Sprintf("%s4", colChar)

				cell4 := fmt.Sprintf("%s5", colChar)

				f.SetCellValue(sheetName, cell1, key.Country)
				f.SetCellValue(sheetName, cell2, key.Oblast)
				f.SetCellValue(sheetName, cell3, key.Region)
				f.SetCellValue(sheetName, cell4, measure)

				f.SetCellStyle(sheetName, cell1, cell1, headerStyle)
				f.SetCellStyle(sheetName, cell2, cell2, headerStyle)
				f.SetCellStyle(sheetName, cell3, cell3, headerStyle)
				f.SetCellStyle(sheetName, cell4, cell4, headerStyle)

				if prevCountry == key.Country && prevCountry != "" {
					prevColChar := service.ToChar(index + len(headers) - 1)
					prevCell1 := fmt.Sprintf("%s2", prevColChar)
					f.MergeCell(sheetName, prevCell1, cell1)
				}
				if prevOblast == key.Oblast && prevOblast != "" {
					prevColChar := service.ToChar(index + len(headers) - 1)
					prevCell1 := fmt.Sprintf("%s3", prevColChar)
					f.MergeCell(sheetName, prevCell1, cell2)
				}
				if prevRegion == key.Region && (prevRegion != "" || (prevRegion == "" && index != 0)) {
					prevColChar := service.ToChar(index + len(headers) - 1)
					prevCell1 := fmt.Sprintf("%s4", prevColChar)
					f.MergeCell(sheetName, prevCell1, cell3)
				}
				if key.Oblast == "" || key.Region == "" {
					f.MergeCell(sheetName, cell2, cell3)
					f.SetCellValue(sheetName, cell2, key.Oblast+key.Region)
				}

				f.SetColWidth(sheetName, "A", colChar, 15)
				index++
				prevCountry = key.Country
				prevOblast = key.Oblast
				prevRegion = key.Region
			}
		}
		productKey := productKey{
			source:          sale.Source,
			barcode:         sale.Barcode,
			name:            sale.Name,
			itemId:          sale.ItemId,
			supplierArticle: sale.SupplierArticle,
			externalCode:    sale.ExternalCode,
		}
		rowIndex, ok := productRow[productKey]
		if !ok {
			row++
			rowIndex = uint32(row)
			productRow[productKey] = rowIndex
		}
		fields := structs.Map(&sale)
		for index, header := range headers {
			colChar := service.ToChar(index)

			cell := fmt.Sprintf("%s%d", colChar, rowIndex+5)
			field := fields[header.field]
			f.SetCellValue(sheetName, cell, field)

			colIndex := int(*customHeader[key])
			cell = fmt.Sprintf("%s%d", service.ToChar(colIndex+len(headers)), rowIndex+5)
			f.SetCellValue(sheetName, cell, sale.Quantity)
			cell = fmt.Sprintf("%s%d", service.ToChar(colIndex+len(headers)+1), rowIndex+5)
			f.SetCellValue(sheetName, cell, sale.TotalPrice)
			cell = fmt.Sprintf("%s%d", service.ToChar(colIndex+len(headers)+2), rowIndex+5)
			f.SetCellValue(sheetName, cell, sale.PriceWithDiscount)
		}
	}

	lastColName := service.ToChar(len(headers) - 1)
	f.AutoFilter(sheetName, fmt.Sprintf("A5:%s5", lastColName), []excelize.AutoFilterOptions{})

	f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s6", service.ToChar(6)),
		Split:       true,
		XSplit:      6,
		YSplit:      5,
		ActivePane:  "bottomLeft",
	})

	return f.Write(writer)
}

func (s *SaleService) GetSale(year uint16, month uint8, page int32, size int32) (*api.Sales, error) {
	sales, err := s.repository.GetSaleByMonthWithPagging(year, month, page, size)
	if err != nil {
		return nil, err
	}
	len := len(*sales)
	if len == 0 {
		return &api.Sales{
			Items: []api.Sale{},
			Count: 0,
		}, nil
	}
	total := int32(0)
	items := make([]api.Sale, len)
	for i, sale := range *sales {
		if total == 0 {
			total = sale.Total
		}
		items[i] = api.Sale{
			Id:              sale.RowNumber,
			Source:          sale.Source,
			Name:            sale.Name,
			SupplierArticle: sale.SupplierArticle,
			Barcode:         sale.Barcode,
			ExternalCode:    sale.ExternalCode,
			Code1c:          sale.ItemId,
			Quantity:        sale.Quantity,
			TotalPrice:      sale.PriceWithDiscount,
			Oblast:          sale.Oblast,
			Region:          sale.Region,
			Country:         sale.Country,
		}
	}
	return &api.Sales{
		Items: items,
		Count: total,
	}, nil
}
