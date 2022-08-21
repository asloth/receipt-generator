package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	//Datos genrales que necesitamos del usuario
	fechaEmision := "22/07/2022"
	fechaVenc := "31/07/2022"
	tipoCuota := "ORDINARIO"
	periodo := "AGOSTO-2022"
	totalPresupuesto := "28,974.00"
	// nameFile := "GPR CUOTA AGOSTO 2022"

	finalColumn := 11
	totalNumberOfRows := 211

	//variable que representa al edificio
	gpr := make(map[string]string)

	gpr["area_total"] = "14,926.79 m2"
	gpr["total_pres"] = totalPresupuesto

	filePath := "cuotas/GPR CUOTA AGOSTO 2022.xlsx"

	sheetName := "Propietarios ordenados"

	fmt.Println("flag1.1")
	xlsxFile, err := excelize.OpenFile(filePath)
	fmt.Println("flag1.2")
	if err != nil {
		fmt.Println(err.Error())
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
	fmt.Println("flag1")
	cols := []string{}

	fmt.Println("flag2")
	ret := []Apartment{}
	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
			fmt.Println("Column information", cols)
		} else {
			ap := Apartment{}
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
	ret[2].GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, gpr["area_total"], gpr["total_pres"])

}
