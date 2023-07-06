package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/yoda/tnot/internal/storage/report"
	"time"
)

func getHeaders() []string {
	return []string{"Кабинет",
		"Площадка",
		"Код склада",
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
		"Текущий остаток товара, шт",
		"Прогноз продаж на 30 дней, шт",
		"В поставке, шт",
		"Нужно поставить, шт",
		"Текущий остаток товара, шт",
		"Кластер",
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
		"Текущий остаток товара, шт",
		"Прогноз продаж на 30 дней, шт",
		"В поставке, шт",
		"Нужно поставить, шт",
		"Текущий остаток товара, шт",
	}
}

func toChar(i int) string {
	return string('A' + i)
}

type reportRepository interface {
	GetReport(date time.Time) (*[]report.Report, error)
	GetReportByCluster(date time.Time) (*[]report.ReportByCluster, error)
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

func (r *ReportService) Print(date time.Time, finaName string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Set value of a cell.
	headers := getHeaders()
	for index, header := range headers {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s2", toChar(index)), header)
	}
	reports, err := r.repository.GetReport(date)
	if err != nil {
		return err
	}
	activeSheet := "Sheet1"
	for index, report := range *reports {
		f.SetCellValue(activeSheet, fmt.Sprintf("A%d", index+3), report.OwnerCode)
		f.SetCellValue(activeSheet, fmt.Sprintf("B%d", index+3), report.Source)
		f.SetCellValue(activeSheet, fmt.Sprintf("C%d", index+3), report.WarehouseName)
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
		f.SetCellValue(activeSheet, fmt.Sprintf("Q%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("R%d", index+3), report.ForecastOrder30)
		f.SetCellValue(activeSheet, fmt.Sprintf("S%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("T%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("U%d", index+3), report.Quantity)
		f.SetCellValue(activeSheet, fmt.Sprintf("V%d", index+3), ifString(report.Cluster))
	}
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет %s`, date.Format("02.01.2006")))

	activeSheet = "Sheet2"
	_, err = f.NewSheet(activeSheet)
	if err != nil {
		return err
	}
	headers = getClusterHeaders()
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
		f.SetCellValue(activeSheet, fmt.Sprintf("Q%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("R%d", index+3), report.ForecastOrder30)
		f.SetCellValue(activeSheet, fmt.Sprintf("S%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("T%d", index+3), "")
		f.SetCellValue(activeSheet, fmt.Sprintf("U%d", index+3), report.Quantity)
	}
	f.SetSheetName(activeSheet, fmt.Sprintf(`Отчет по кластерам %s`, date.Format("02.01.2006")))

	// Set active sheet of the workbook.
	// Save spreadsheet by the given path.
	return f.SaveAs(finaName)
}
