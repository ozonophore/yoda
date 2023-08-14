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

func toChar(i int) string {
	return string('A' + i)
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

func drawHeaders(sheet string, headers []column, f *excelize.File) {
	headerStyle, _ := f.NewStyle(&excelize.Style{
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
	})
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
		fmt.Println(err)
	}

	// Set active sheet of the workbook.
	// Save spreadsheet by the given path.
	return f.SaveAs(finaName)
}
