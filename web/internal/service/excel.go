package service

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/xuri/excelize/v2"
	"io"
	"reflect"
	"strings"
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
	value := string('A' + rune(i))
	if i > 25 {
		pref := string('A' + rune(i/26-1))
		value = pref + string('A'+rune(i%26))
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

func GetExcelHeaders(model interface{}) *[]ExcelHeaderColumn {
	typeOf := reflect.TypeOf(model)

	list := make([]ExcelHeaderColumn, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		excelTag := field.Tag.Get("excel")
		if excelTag == "" {
			continue
		}
		header := parseMetadata(excelTag)
		header.Field = field.Name
		list = append(list, header)
	}
	return &list
}

func parseMetadata(tag string) ExcelHeaderColumn {
	metadata := ExcelHeaderColumn{}

	// Разделение строки метаданных по точкам с запятыми
	pairs := strings.Split(tag, ";")

	// Обработка каждой пары ключ-значение
	for _, pair := range pairs {
		// Разделение пары по двоеточиям
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			value := strings.TrimSpace(parts[1])

			switch key {
			case "title":
				metadata.Title = value
			case "width":
				var width float64
				fmt.Sscanf(value, "%f", &width)
				metadata.Width = width
			}
		}
	}

	return metadata
}
