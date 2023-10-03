package service

import "github.com/xuri/excelize/v2"

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
