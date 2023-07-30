package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/yoda/tnot/internal/storage/report"
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
		{title: "Продажи за 5 дней неделю назад, шт", name: "", width: 15},
		{title: "Дней в дефектуре за 30 дней", name: "def30", width: 15},
		{title: "Дней в дефектуре за 5 дней", name: "def5", width: 15},
		{title: "Прогноз продаж на 30 дней, шт", name: "forecast_order30", width: 15},
		{title: "Прогноз продаж на 5 дней, шт", name: "forecast_order5", width: 15},
		{title: "Прогноз продаж средний, шт", name: "forecast_avg", width: 15},
		{title: "В поставке, шт", name: "", width: 15},
		{title: "Нужно поставить, шт", name: "", width: 15},
		{title: "Остаток склада общий, шт", name: "quantity1c", width: 15},
		{title: "Остаток склада в % по складам", name: "stock1c_percent", width: 15},
	}
}

func getClusterHeaders() []string {
	return []string{"Кабинет",
		"Площадка",
		"Кластер",
		"ID товара",
		"Наименование",
		"Артикул",
		"Артикул 1С",
		"Наименование",
		"Штрихкод",
		"Выведенная позиция",
		"Комплект, шт",
		"Продажи за 30 дней, шт",
		"Продажи за 5 дней, шт",
		"Продажи за 5 дней неделю назад, шт",
		"Дней в дефектуре за 30 дней",
		"Дней в дефектуре за 5 дней",
		"Прогноз продаж на 30 дней, шт",
		"В поставке, шт",
		"Нужно поставить, шт",
		"Текущий остаток товара, шт",
	}
}

func getItemHeaders() []string {
	return []string{"Кабинет",
		"Площадка",
		"ID товара",
		"Наименование",
		"Артикул",
		"Артикул 1С",
		"Штрихкод",
		"Выведенная позиция",
		"Комплект, шт",
		"Продажи за 30 дней, шт",
		"Продажи за 5 дней, шт",
		"Продажи за 5 дней неделю назад, шт",
		"Дней в дефектуре за 30 дней",
		"Дней в дефектуре за 5 дней",
		"Прогноз продаж на 30 дней, шт",
		"В поставке, шт",
		"Нужно поставить, шт",
		"Текущий остаток товара, шт",
		"Текущий остаток из 1С товара, шт",
	}
}

func toChar(i int) string {
	return string('A' + i)
}

type reportRepository interface {
	GetReport(date time.Time) (*[]map[string]interface{}, error)
	GetReportByCluster(date time.Time) (*[]report.ReportByCluster, error)
	GetReportByItem(date time.Time) (*[]report.ReportByItem, error)
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

func drawHeaders(headers []column, f *excelize.File) {
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
		f.SetCellValue("Sheet1", cell, header.title)
		f.SetCellStyle("Sheet1", cell, cell, headerStyle)

		f.SetColWidth("Sheet1", "A", toChar(index), header.width)
	}
	lastColName := toChar(len(headers) - 1)
	f.AutoFilter("Sheet1", fmt.Sprintf("A2:%s2", lastColName), []excelize.AutoFilterOptions{})
	f.SetPanes("Sheet1", &excelize.Panes{
		Freeze:      true,
		TopLeftCell: fmt.Sprintf("%s3", toChar(9)),
		Split:       true,
		XSplit:      9,
		YSplit:      2,
		ActivePane:  "bottomLeft",
	})
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
	drawHeaders(headersByWarehouses, f)

	//Fill the body
	reports, err := r.repository.GetReport(date)
	if err != nil {
		return err
	}
	activeSheet := "Sheet1"
	for index, report := range *reports {
		for i, header := range headersByWarehouses {
			colName := toChar(i)
			value := report[header.name]
			if _, ok := value.(bool); ok {
				value = ifBool(value.(bool))
			}
			f.SetCellValue(activeSheet, fmt.Sprintf("%s%d", colName, index+3), value)
		}
	}
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет %s`, date.Format("02.01.2006")))

	activeSheet = "Sheet2"
	_, err = f.NewSheet(activeSheet)
	if err != nil {
		return err
	}
	headers := getClusterHeaders()
	for index, header := range headers {
		f.SetCellValue(activeSheet, fmt.Sprintf("%s2", toChar(index)), header)
	}
	clusters, err := r.repository.GetReportByCluster(date)
	if err != nil {
		return err
	}
	for index, report := range *clusters {
		f.SetCellValue(activeSheet, fmt.Sprintf("A%d", index+3), report.OwnerCode)
		f.SetCellValue(activeSheet, fmt.Sprintf("B%d", index+3), report.Source)
		f.SetCellValue(activeSheet, fmt.Sprintf("C%d", index+3), ifString(report.Cluster))
		f.SetCellValue(activeSheet, fmt.Sprintf("D%d", index+3), report.ExternalCode)
		f.SetCellValue(activeSheet, fmt.Sprintf("E%d", index+3), report.ItemName)
		f.SetCellValue(activeSheet, fmt.Sprintf("F%d", index+3), report.SupplierArticle)
		f.SetCellValue(activeSheet, fmt.Sprintf("G%d", index+3), report.ItemId)
		f.SetCellValue(activeSheet, fmt.Sprintf("H%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("I%d", index+3), report.Barcode)
		f.SetCellValue(activeSheet, fmt.Sprintf("J%d", index+3), ifBool(report.IsExcluded))
		f.SetCellValue(activeSheet, fmt.Sprintf("K%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("L%d", index+3), report.Quantity30)
		f.SetCellValue(activeSheet, fmt.Sprintf("M%d", index+3), report.Quantity5)
		f.SetCellValue(activeSheet, fmt.Sprintf("N%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("O%d", index+3), report.Def30)
		f.SetCellValue(activeSheet, fmt.Sprintf("P%d", index+3), report.Def5)
		f.SetCellValue(activeSheet, fmt.Sprintf("Q%d", index+3), report.ForecastOrder30)
		f.SetCellValue(activeSheet, fmt.Sprintf("R%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("S%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("T%d", index+3), report.Quantity)
	}
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет по кластерам %s`, date.Format("02.01.2006")))

	sheetIndex, err := f.NewSheet("Сводный отчет")
	if err != nil {
		return err
	}
	sheetName := f.GetSheetName(sheetIndex)
	headers = getItemHeaders()
	for index, header := range headers {
		f.SetCellValue(sheetName, fmt.Sprintf("%s2", toChar(index)), header)
	}
	items, err := r.repository.GetReportByItem(date)
	if err != nil {
		return err
	}
	for index, item := range *items {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", index+3), item.OwnerCode)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", index+3), item.Source)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", index+3), item.ExternalCode)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", index+3), item.ItemName)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", index+3), item.SupplierArticle)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", index+3), item.ItemId)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", index+3), item.Barcode)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", index+3), ifBool(item.IsExcluded))

		f.SetCellValue(sheetName, fmt.Sprintf("J%d", index+3), item.Quantity30)
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", index+3), item.Quantity5)

		f.SetCellValue(sheetName, fmt.Sprintf("M%d", index+3), item.Def30)
		f.SetCellValue(sheetName, fmt.Sprintf("N%d", index+3), item.Def5)
		f.SetCellValue(sheetName, fmt.Sprintf("O%d", index+3), item.ForecastOrder30)

		f.SetCellValue(sheetName, fmt.Sprintf("R%d", index+3), item.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("S%d", index+3), item.Quantity1С)
	}
	f.AutoFilter(sheetName, "A2:R2", []excelize.AutoFilterOptions{})
	err = f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		TopLeftCell: "F3",
		YSplit:      2,
		XSplit:      5,
		ActivePane:  "bottomLeft",
	})
	if err != nil {
		fmt.Println(err)
	}

	// Set active sheet of the workbook.
	// Save spreadsheet by the given path.
	return f.SaveAs(finaName)
}
