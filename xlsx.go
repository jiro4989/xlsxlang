package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ReadXlsx(path string) (string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return "", err
	}

	var result string
	for y, row := range rows {
		if y == 0 {
			continue
		}
		for x, cell := range row {
			if x == 0 {
				continue
			}
			if cell == "" {
				cell = " "
			}
			result += cell
		}
		result += "\n"
	}
	return result, nil
}
