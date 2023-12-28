package service

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/xuri/excelize/v2"
	"io"
)

type ExcelHeaderColumn struct {
	Title string
	Width float64
	Field string
}

func GetHeaderStyle() *excelize.Style {
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

func ToChar(i int) string {
	value := string('A' + i)
	if i > 25 {
		pref := string('A' + i/26 - 1)
		value = pref + string('A'+i%26)
	}
	return value
}

func createHeaders(f *excelize.File, sheetName string, headerStyle int, headers *[]ExcelHeaderColumn, rowIndex int) {
	for index, header := range *headers {
		colChar := ToChar(index)
		cell := fmt.Sprintf("%s%d", colChar, rowIndex)
		f.SetCellValue(sheetName, cell, header.Title)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
		f.SetColWidth(sheetName, "A", colChar, header.Width)
	}
}

func GenerateExcelDoc[T any](writer io.Writer, sheetName string, data *[]T, headers *[]ExcelHeaderColumn) error {
	f := excelize.NewFile()
	defer f.Close()

	headerStyle, _ := f.NewStyle(GetHeaderStyle())
	sheetIndex := f.GetActiveSheetIndex()

	f.SetSheetName(f.GetSheetName(sheetIndex), sheetName)

	rowIndex := 3
	for index, item := range *data {
		if index == 0 {
			createHeaders(f, sheetName, headerStyle, headers, rowIndex)
		}
		fields := structs.Map(&item)
		for i, header := range *headers {
			colChar := ToChar(i)
			cell := fmt.Sprintf("%s%d", colChar, rowIndex+index+1)
			field := fields[header.Field]
			f.SetCellValue(sheetName, cell, field)
		}
	}

	lastColName := ToChar(len(*headers) - 1)
	f.AutoFilter(sheetName, fmt.Sprintf("A%d:%s%d", rowIndex, lastColName, rowIndex), []excelize.AutoFilterOptions{})

	return f.Write(writer)
}
