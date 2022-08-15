package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	// fechaEmision := "22/07/2022"
	// fechaVenc := "31/07/2022"
	// tipoCuota := "ORDINARIO"
	finalColumn := 11
	// startColumn := "A"
	totalNumberOfRows := 210

	filePath := "files/GPR-CUOTA-JULIO-2022.xlsx"

	sheetName := "Propietarios ordenados"

	xlsxFile, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := xlsxFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := xlsxFile.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	cols := []string{}
	ret := []Apartment{}
	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
			fmt.Println("Column information", cols)
		} else {
			var ap Apartment
			for j, colCell := range row {
				if j > finalColumn {
					break
				}

				switch cols[j] {
				case "propietario":
					ap.owner = colCell
				case "depa":
					ap.number, _ = strconv.ParseInt(colCell, 10, 64)
				case "total área":
					ap.totalArea, _ = strconv.ParseFloat(colCell, 64)
				case "cuota":
					ap.amount, _ = strconv.ParseFloat(colCell, 64)
				case "porcentaje":
					ap.percentaje, _ = strconv.ParseFloat(colCell, 64)
				case "estaciona":
					if len(colCell) == 0 {
						colCell = "--"
					}
					ap.parking = colCell
				}
			}
			ret = append(ret, ap)
		}
		if i > totalNumberOfRows {
			break
		}

	}
	fmt.Println(ret[2])

}

type Apartment struct {
	number     int64
	owner      string
	totalArea  float64
	percentaje float64
	amount     float64
	parking    string
}
