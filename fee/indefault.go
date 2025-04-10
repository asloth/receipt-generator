package fee

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type InDefault struct {
	ApartmentNumber string
	Amount          float64
}

func LoadInDefaultData(filePath, sheetName string) ([]InDefault, error) {

	xlsxFile, err := excelize.OpenFile(filePath)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := xlsxFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the
	rows, err := xlsxFile.GetRows(sheetName)

	if err != nil {
		return nil, err
	}

	ret := []InDefault{}

	for i, row := range rows {
		if i == 0 {
			continue
		} else {
			ap := InDefault{}
		inside:
			for j, colCell := range row {
				colCell = strings.TrimSpace(colCell) //el valor de la celda
				if j == 0 {
					ap.ApartmentNumber = colCell
					continue inside
				}
				ap.Amount, err = strconv.ParseFloat(colCell, 64)
				if err != nil {
					return nil, err
				}
			}
			ret = append(ret, ap)
		}

	}
	return ret, nil
}
