package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GENERAR RECIBOS")
	fmt.Println("---------------------")

	//Datos genrales que necesitamos del usuario
	fechaEmision := ""
	getReceiptData(reader, "fecha de emision (dd/mm/aa)", &fechaEmision, true)

	fechaVenc := ""
	getReceiptData(reader, "fecha de vencimiento (dd/mm/aa)", &fechaVenc, true)

	tipoCuota := "ORDINARIO"
	getReceiptData(reader, "tipo de cuota", &tipoCuota, false)

	periodo := "AGOSTO-2022"
	getReceiptData(reader, "periodo", &periodo, false)

	totalPresupuesto := "28,974.00"
	getFloatData(reader, "presupuesto", &totalPresupuesto)

	// Limits in the spreadsheet
	finalColumn := 13
	totalNumberOfRows := 212

	//variable que representa al edificio
	gpr := make(map[string]string)

	gpr["total_pres"] = totalPresupuesto

	filePath := "cuotas/GPR CUOTA SETIEMBRE 2022.xlsx"

	sheetName := "Propietarios ordenados"

	// Open the spreadsheet
	xlsxFile, err := excelize.OpenFile(filePath)

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

	cols := []string{}

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

				switch strings.ToLower(cols[j]) {
				case "propietario":
					if len(colCell) == 0 {
						colCell = "Sin datos"
					}
					ap.owner = colCell
				case "depa":
					ap.number, err = strconv.ParseInt(colCell, 10, 64)
					if err != nil {
						ap.number = 0.0
					}
				case "total área":
					ap.totalArea, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.number = 0.0
					}
				case "área-e":
					ap.parkingArea, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.number = 0.0
					}
				case "total":
					ap.total, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.total = 0.0
					}
				case "cuota":
					ap.maintenance, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.maintenance = 0.0
					}
				case "porcentaje":
					ap.percentaje, _ = strconv.ParseFloat(colCell, 64)
				case "estaciona":
					if len(colCell) == 0 {
						colCell = "--"
					}
					ap.parking = colCell
				case "agua":
					ap.waterComsuption, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.waterComsuption = 0.0
					}
				case "deposito":
					if len(colCell) == 0 {
						colCell = "--"
					}
					ap.deposit = append(ap.deposit, colCell)
				default:
					continue

				}
			}
			ret = append(ret, ap)
		}
		if i > totalNumberOfRows {
			break
		}

	}
	fmt.Println(ret[210])
	ret[210].GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, gpr["total_pres"])

}

func getReceiptData(r *bufio.Reader, question string, data *string, isADate bool) {
	reader := *r
	for {
		// Ask for the date
		fmt.Print("Ingresar " + question + ": ")

		// Reading the user input
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		stringDate := strings.Replace(text, "\n", "", -1)

		if isADate {
			// Verify the date's format
			_, err := time.Parse("02/01/2006", stringDate)

			if err != nil {
				// Show error
				fmt.Println("ERROR: Dato invalido")
				fmt.Println(err)

				// Ask again if the date if not correct
				continue
			}
		}

		*data = stringDate
		fmt.Println(question + " guardada: " + *data)

		break

	}
}

func getFloatData(r *bufio.Reader, question string, data *string) {
	reader := *r
	for {
		// Ask for the date
		fmt.Print("Ingresar " + question + ": ")

		// Reading the user input
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input := strings.Replace(text, "\n", "", -1)

		if _, err := strconv.ParseFloat(input, 64); err != nil {
			fmt.Printf("ERROR: Dato ingresado no es un numero" + input)
			continue
		}

		*data = input
		fmt.Println(question + " guardada: " + *data)

		break

	}
}
