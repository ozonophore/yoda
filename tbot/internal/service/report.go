package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

type column struct {
	title string
	name  string
	width float64
}

type groupColumn struct {
	group string
	title string
	name  string
	width float64
	color string
}

func getHeaders() []column {
	return []column{
		{title: "Кабинет", name: "owner_code", width: 14},
		{title: "Площадка", name: "source", width: 15},
		{title: "Кластер", name: "cluster", width: 15},
		{title: "Код склада", name: "warehouse_name", width: 20},
		{title: "ID товара", name: "external_code", width: 15},
		{title: "Наименование", name: "item_name", width: 25},
		{title: "Артикул", name: "supplier_article", width: 20},
		{title: "Артикул 1С", name: "item_id", width: 15},
		{title: "Штрихкод", name: "barcode", width: 20},
		{title: "Выведенная позиция", name: "is_excluded", width: 15},
		{title: "Комплект, шт", name: "", width: 15},
		{title: "Текущий остаток товара, шт", name: "quantity", width: 15},
		{title: "Продажи за 30 дней, шт", name: "quantity30", width: 15},
		{title: "Продажи за 5 дней, шт", name: "quantity5", width: 15},
		{title: "Продажи за 5 дней неделю назад, шт", name: "quantity5_week_ago", width: 15},
		{title: "Дней в дефектуре за 30 дней", name: "def30", width: 15},
		{title: "Дней в дефектуре за 5 дней", name: "def5", width: 15},
		{title: "Оборачивоемость, 30 дн", name: "turnover30", width: 15},
		{title: "Оборачивоемость, 5 дн", name: "turnover5", width: 15},
		{title: "Прогноз продаж на 30 дней, шт", name: "forecast_order30", width: 15},
		{title: "Прогноз продаж на 5 дней, шт", name: "forecast_order5", width: 15},
		{title: "Прогноз продаж средний, шт", name: "forecast_avg", width: 15},
		{title: "В поставке, шт", name: "", width: 15},
		{title: "Нужно поставить, шт", name: "", width: 15},
		{title: "Остаток склада общий, шт", name: "quantity1c", width: 15},
		{title: "Остаток склада в % по складам", name: "stock1c_percent", width: 15},
	}
}

func getClusterHeaders() []column {
	return []column{
		{title: "Кабинет", name: "owner_code", width: 14},
		{title: "Площадка", name: "source", width: 15},
		{title: "Кластер", name: "cluster", width: 15},
		{title: "ID товара", name: "external_code", width: 15},
		{title: "Наименование", name: "item_name", width: 25},
		{title: "Артикул", name: "supplier_article", width: 20},
		{title: "Код 1С", name: "item_id", width: 15},
		{title: "Сегмент", name: "segment", width: 10},
		{title: "Бренд", name: "brand", width: 20},
		{title: "РРЦ", name: "retail_price", width: 15},
		{title: "Штрихкод", name: "barcode", width: 20},
		{title: "Выведенная позиция", name: "is_excluded", width: 15},
		{title: "Комплект, шт", name: "", width: 15},
		{title: "Текущий остаток товара, шт", name: "quantity", width: 15},
		{title: "Продажи за 30 дней, шт", name: "quantity30", width: 15},
		{title: "Продажи за 5 дней, шт", name: "quantity5", width: 15},
		{title: "Продажи за 5 дней неделю назад, шт", name: "quantity5_week_ago", width: 15},
		{title: "Дней в дефектуре за 30 дней", name: "def30", width: 15},
		{title: "Дней в дефектуре за 5 дней", name: "def5", width: 15},
		{title: "Оборачивоемость, 30 дн", name: "turnover30", width: 15},
		{title: "Оборачивоемость, 5 дн", name: "turnover5", width: 15},
		{title: "Прогноз продаж на 30 дней, шт", name: "forecast_order30", width: 15},
		{title: "Прогноз продаж на 5 дней, шт", name: "forecast_order5", width: 15},
		{title: "В поставке, шт", name: "", width: 15},
		{title: "Нужно поставить, шт", name: "", width: 15},
		{title: "Остаток склада общий, шт", name: "quantity1c", width: 15},
		{title: "Остаток склада в % по кластерам", name: "stock1c_percent", width: 15},
	}
}

func getItemHeaders() []column {
	return []column{
		{title: "Кабинет", name: "owner_code", width: 14},
		{title: "Площадка", name: "source", width: 15},
		{title: "ID товара", name: "external_code", width: 15},
		{title: "Наименование", name: "item_name", width: 25},
		{title: "Артикул", name: "supplier_article", width: 20},
		{title: "Артикул 1С", name: "item_id", width: 15},
		{title: "Штрихкод", name: "barcode", width: 20},
		{title: "Выведенная позиция", name: "is_excluded", width: 15},
		{title: "Комплект, шт", name: "", width: 15},
		{title: "Текущий остаток товара, шт", name: "quantity", width: 15},
		{title: "Продажи за 30 дней, шт", name: "quantity30", width: 15},
		{title: "Продажи за 5 дней, шт", name: "quantity5", width: 15},
		{title: "Продажи за 5 дней неделю назад, шт", name: "quantity5_week_ago", width: 15},
		{title: "Дней в дефектуре за 30 дней", name: "def30", width: 15},
		{title: "Дней в дефектуре за 5 дней", name: "def5", width: 15},
		{title: "Прогноз продаж на 30 дней, шт", name: "forecast_order30", width: 15},
		{title: "Прогноз продаж на 5 дней, шт", name: "forecast_order5", width: 15},
		{title: "Оборачивоемость, 30 дн", name: "turnover30", width: 15},
		{title: "Оборачивоемость, 5 дн", name: "turnover5", width: 15},
		{title: "В поставке, шт", name: "", width: 15},
		{title: "Нужно поставить, шт", name: "", width: 15},
		{title: "Текущий остаток товара, шт", name: "quantity", width: 15},
		{title: "Текущий остаток из 1С товара, шт", name: "quantity1c", width: 15},
		{title: "Остаток склада в % по кластерам", name: "stock1c_percent", width: 15},
	}
}

func getPositionHeader() []groupColumn {
	return []groupColumn{
		{
			group: "",
			title: "Выведенная позиция",
			name:  "total_is_excluded",
			width: 15,
			color: "",
		}, {
			group: "",
			title: "Код 1С",
			name:  "total_item_id",
			width: 15,
			color: "",
		},
		{
			group: "",
			title: "Сегмент",
			name:  "segment",
			width: 15,
			color: "",
		}, {
			group: "",
			title: "Бренд",
			name:  "brand",
			width: 20,
			color: "",
		}, {
			group: "",
			title: "РРЦ",
			name:  "retail_price",
			width: 15,
			color: "",
		},
		{
			group: "",
			title: "Наименование",
			name:  "total_item_name",
			width: 25,
			color: "",
		}, {
			group: "",
			title: "Комплект, шт",
			name:  "",
			width: 15,
			color: "",
		}, {
			group: "OZON",
			title: "Текущий остаток товара, шт",
			name:  "ozon_quantity",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Продажи за 30 дней, шт",
			name:  "ozon_quantity30",
			width: 15,
			color: "E4EFDC",
		},
		{
			group: "OZON",
			title: "Продажи за 5 дней, шт",
			name:  "ozon_quantity5",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Дней в дефектуре за 30 дней",
			name:  "ozon_def30",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Дней в дефектуре за 5 дней",
			name:  "ozon_def5",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "ozon_forecast_order30",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "ozon_forecast_order5",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Прогноз ср",
			name:  "ozon_forecast_avg",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Оборачиваемость 30 дн",
			name:  "ozon_turnover30",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "OZON",
			title: "Оборачиваемость 5 дн",
			name:  "ozon_turnover5",
			width: 15,
			color: "E4EFDC",
		}, {
			group: "WB",
			title: "Текущий остаток товара, шт",
			name:  "wb_quantity",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Продажи за 30 дней, шт",
			name:  "wb_quantity30",
			width: 15,
			color: "FDF2D0",
		},
		{
			group: "WB",
			title: "Продажи за 5 дней, шт",
			name:  "wb_quantity5",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Дней в дефектуре за 30 дней",
			name:  "wb_def30",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Дней в дефектуре за 5 дней",
			name:  "wb_def5",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "wb_forecast_order30",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "wb_forecast_order5",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Прогноз ср",
			name:  "wb_forecast_avg",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Оборачиваемость 30 дн",
			name:  "wb_turnover30",
			width: 15,
			color: "FDF2D0",
		}, {
			group: "WB",
			title: "Оборачиваемость 5 дн",
			name:  "wb_turnover5",
			width: 15,
			color: "FDF2D0",
		},
		{
			group: "Total",
			title: "Нужно поставить, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Текущий остаток из 1С товара, шт",
			name:  "total_quantity1c",
			width: 15,
		},
		{
			group: "Total",
			title: "Текущий остаток товара, шт",
			name:  "total_quantity",
			width: 15,
		}, {
			group: "Total",
			title: "Продажи за 30 дней, шт",
			name:  "total_quantity30",
			width: 15,
		},
		{
			group: "Total",
			title: "Продажи за 5 дней, шт",
			name:  "total_quantity5",
			width: 15,
		}, {
			group: "Total",
			title: "Дней в дефектуре за 30 дней",
			name:  "total_def30",
			width: 15,
		}, {
			group: "Total",
			title: "Дней в дефектуре за 5 дней",
			name:  "total_def5",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "total_forecast_order30",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "total_forecast_order5",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз ср",
			name:  "total_forecast_avg",
			width: 15,
		}, {
			group: "Total",
			title: "Оборачиваемость 30 дн",
			name:  "total_turnover30",
			width: 15,
		}, {
			group: "Total",
			title: "Оборачиваемость 5 дн",
			name:  "total_turnover5",
			width: 15,
		}, {
			group: "Total",
			title: "Нужно поставить, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Текущий остаток из 1С товара, шт",
			name:  "total_quantity1c",
			width: 18,
		}, {
			group: "Total",
			title: "Оборачиваемость склада ликато",
			name:  "",
			width: 18,
		}, {
			group: "Total",
			title: "Прогнозная Оборачиваемость склада ликато",
			name:  "",
			width: 20,
		},
	}
}

func getRoomHeader() []groupColumn {
	return []groupColumn{
		{
			group: "",
			title: "Выведенная позиция",
			name:  "total_is_excluded",
			width: 15,
		}, {
			group: "",
			title: "Артикул 1С",
			name:  "",
			width: 15,
		},
		{
			group: "",
			title: "Наименование",
			name:  "",
			width: 25,
		},
		{
			group: "",
			title: "Кабинет",
			name:  "",
			width: 25,
		},
		{
			group: "",
			title: "Комплект, шт",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Текущий остаток товара, шт",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Продажи за 30 дней, шт",
			name:  "",
			width: 15,
		},
		{
			group: "OZON",
			title: "Продажи за 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Дней в дефектуре за 30 дней",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Дней в дефектуре за 5 дней",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Прогноз ср",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Оборачиваемость 30 дн",
			name:  "",
			width: 15,
		}, {
			group: "OZON",
			title: "Оборачиваемость 5 дн",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Текущий остаток товара, шт",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Продажи за 30 дней, шт",
			name:  "",
			width: 15,
		},
		{
			group: "WB",
			title: "Продажи за 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Дней в дефектуре за 30 дней",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Дней в дефектуре за 5 дней",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Прогноз ср",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Оборачиваемость 30 дн",
			name:  "",
			width: 15,
		}, {
			group: "WB",
			title: "Оборачиваемость 5 дн",
			name:  "",
			width: 15,
		},
		{
			group: "Total",
			title: " Нужно поставить, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: " Текущий остаток из 1С товара, шт",
			name:  "",
			width: 15,
		},
		{
			group: "Total",
			title: "Текущий остаток товара, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Продажи за 30 дней, шт",
			name:  "",
			width: 15,
		},
		{
			group: "Total",
			title: "Продажи за 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Дней в дефектуре за 30 дней",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Дней в дефектуре за 5 дней",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз продаж на 30 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз продаж на 5 дней, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Прогноз ср",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Оборачиваемость 30 дн",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Оборачиваемость 5 дн",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Нужно поставить, шт",
			name:  "",
			width: 15,
		}, {
			group: "Total",
			title: "Текущий остаток из 1С товара, шт",
			name:  "",
			width: 18,
		}, {
			group: "Total",
			title: "Оборачиваемость склада ликато",
			name:  "",
			width: 18,
		}, {
			group: "Total",
			title: "Прогнозная Оборачиваемость склада ликато",
			name:  "",
			width: 20,
		},
	}
}

func toChar(i int) string {
	value := string('A' + i)
	if i > 25 {
		value = "A" + string('A'+i-26)
	}
	return value
}

type reportRepository interface {
	GetReportByProduct(date time.Time) (*[]map[string]interface{}, error)
	GetReport(date time.Time) (*[]map[string]interface{}, error)
	GetReportByCluster(date time.Time) (*[]map[string]interface{}, error)
	GetReportByItem(date time.Time) (*[]map[string]interface{}, error)
}

type ReportService struct {
	repository reportRepository
}

func NewReportService(repository reportRepository) *ReportService {
	return &ReportService{repository: repository}
}

func ifBool(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}

func ifString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getHeaderStyle() *excelize.Style {
	return &excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Alignment: &excelize.Alignment{
			WrapText:    true,
			ShrinkToFit: true,
			Horizontal:  "center",
			Vertical:    "center",
		},
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"99CCFF"}},
		Border: []excelize.Border{
			{Type: "left", Color: "808080", Style: 1},
			{Type: "top", Color: "808080", Style: 1},
			{Type: "bottom", Color: "808080", Style: 1},
			{Type: "right", Color: "808080", Style: 1},
		},
	}
}

func getCellStyle(color string) *excelize.Style {
	return &excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{color}},
		Border: []excelize.Border{
			{Type: "left", Color: "808080", Style: 1},
			{Type: "top", Color: "808080", Style: 1},
			{Type: "bottom", Color: "808080", Style: 1},
			{Type: "right", Color: "808080", Style: 1},
		},
	}
}

func drawGroupHeaders(sheet string, headers []groupColumn, f *excelize.File) {
	headerStyle, _ := f.NewStyle(getHeaderStyle())

	prevGroup := ""
	for index, header := range headers {
		cell := fmt.Sprintf("%s3", toChar(index))
		f.SetCellValue(sheet, cell, header.title)
		f.SetCellStyle(sheet, cell, cell, headerStyle)

		f.SetColWidth(sheet, "A", toChar(index), header.width)

		if header.group != "" {
			f.SetCellValue(sheet, fmt.Sprintf("%s2", toChar(index)), header.group)
			f.SetCellStyle(sheet, fmt.Sprintf("%s2", toChar(index)), fmt.Sprintf("%s2", toChar(index)), headerStyle)
		}
		if prevGroup != "" && prevGroup == header.group {
			f.MergeCell(sheet, fmt.Sprintf("%s2", toChar(index-1)), fmt.Sprintf("%s2", toChar(index)))
		}
		prevGroup = header.group
	}
	lastColName := toChar(len(headers) - 1)
	f.AutoFilter(sheet, fmt.Sprintf("A3:%s3", lastColName), []excelize.AutoFilterOptions{})
}

func drawHeaders(sheet string, headers []column, f *excelize.File) {
	headerStyle, _ := f.NewStyle(getHeaderStyle())
	for index, header := range headers {
		cell := fmt.Sprintf("%s2", toChar(index))
		f.SetCellValue(sheet, cell, header.title)
		f.SetCellStyle(sheet, cell, cell, headerStyle)

		f.SetColWidth(sheet, "A", toChar(index), header.width)
	}
	lastColName := toChar(len(headers) - 1)
	f.AutoFilter(sheet, fmt.Sprintf("A2:%s2", lastColName), []excelize.AutoFilterOptions{})
}

func drawBody(sheet string, header *[]column, data *[]map[string]interface{}, f *excelize.File) {
	for index, report := range *data {
		for i, header := range *header {
			colName := toChar(i)
			value := report[header.name]
			if _, ok := value.(bool); ok {
				value = ifBool(value.(bool))
			}
			f.SetCellValue(sheet, fmt.Sprintf("%s%d", colName, index+3), value)
		}
	}
}

func drawBodyWithGroup(sheet string, header *[]groupColumn, data *[]map[string]interface{}, f *excelize.File) {
	for index, report := range *data {
		for i, header := range *header {
			colName := toChar(i)
			value := report[header.name]
			if _, ok := value.(bool); ok {
				value = ifBool(value.(bool))
			}
			cell := fmt.Sprintf("%s%d", colName, index+4)
			f.SetCellValue(sheet, cell, value)
			if header.color != "" {
				styleId, _ := f.NewStyle(getCellStyle(header.color))
				f.SetCellStyle(sheet, cell, cell, styleId)
			}
		}
	}
}

func (r *ReportService) Print(date time.Time, finaName string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headersByWarehouses := getHeaders()
	// Draw  the header
	activeSheet := "Sheet1"
	drawHeaders(activeSheet, headersByWarehouses, f)
	f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s3", toChar(9)),
		Split:       true,
		XSplit:      9,
		YSplit:      2,
		ActivePane:  "bottomLeft",
	})

	//Fill the body
	reports, err := r.repository.GetReport(date)
	if err != nil {
		return err
	}
	drawBody(activeSheet, &headersByWarehouses, reports, f)
	f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s3", toChar(9)),
		Split:       true,
		XSplit:      9,
		YSplit:      2,
		ActivePane:  "bottomLeft",
	})
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет %s`, date.Format("02.01.2006")))

	// Draw the second sheet
	sheetIndex, err := f.NewSheet(activeSheet)
	if err != nil {
		return err
	}
	activeSheet = f.GetSheetName(sheetIndex)

	headers := getClusterHeaders()
	drawHeaders(activeSheet, headers, f)
	clusters, err := r.repository.GetReportByCluster(date)
	if err != nil {
		return err
	}
	drawBody(activeSheet, &headers, clusters, f)

	f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s3", toChar(9)),
		Split:       true,
		XSplit:      9,
		YSplit:      2,
		ActivePane:  "bottomLeft",
	})
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет по кластерам %s`, date.Format("02.01.2006")))

	sheetIndex, err = f.NewSheet("Сводный отчет")
	if err != nil {
		return err
	}
	activeSheet = f.GetSheetName(sheetIndex)
	headers = getItemHeaders()
	drawHeaders(activeSheet, headers, f)
	items, err := r.repository.GetReportByItem(date)
	if err != nil {
		return err
	}
	drawBody(activeSheet, &headers, items, f)

	err = f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: "I3",
		YSplit:      2,
		XSplit:      8,
		ActivePane:  "bottomLeft",
	})
	if err != nil {
		return err
	}
	//--------------------
	sheetIndex, err = f.NewSheet("На уровне позиции")
	if err != nil {
		return err
	}
	activeSheet = f.GetSheetName(sheetIndex)
	groupHeaders := getPositionHeader()
	drawGroupHeaders(activeSheet, groupHeaders, f)
	items, err = r.repository.GetReportByProduct(date)
	drawBodyWithGroup(activeSheet, &groupHeaders, items, f)

	f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s4", toChar(4)),
		Split:       true,
		XSplit:      4,
		YSplit:      3,
		ActivePane:  "bottomLeft",
	})
	//-------------------
	sheetIndex, err = f.NewSheet("На уровне кабинетов")
	if err != nil {
		return err
	}
	activeSheet = f.GetSheetName(sheetIndex)
	groupHeaders = getRoomHeader()
	drawGroupHeaders(activeSheet, groupHeaders, f)

	f.SetPanes(activeSheet, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s4", toChar(5)),
		Split:       true,
		XSplit:      5,
		YSplit:      3,
		ActivePane:  "bottomLeft",
	})
	// Set active sheet of the workbook.
	// Save spreadsheet by the given path.
	return f.SaveAs(finaName)
}
