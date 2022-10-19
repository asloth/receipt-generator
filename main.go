package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	// "os"
	"strconv"
	"strings"
	"time"

	"github.com/asloth/receipt-generator/email"
	"github.com/xuri/excelize/v2"
)

func main2() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GENERAR RECIBOS")
	fmt.Println("---------------------")

	// Datos genrales que necesitamos del usuario
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

	waterRead := ""
	getReceiptData(reader, "fecha de lectura del agua (dd/mm/aa)", &waterRead, true)
	// Limits in the spreadsheet

	// variable que representa al edificio
	gpr := make(map[string]string)

	gpr["total_pres"] = totalPresupuesto

	filePath := "cuotas/GPR CUOTA OCTUBRE 2022.xlsx"

	sheetName := "Propietarios ordenados"

	ret, err := loadApartmentData(filePath, sheetName)

	if err != nil {
		fmt.Println("Error reading apartment data" + err.Error())
	}

	waterData, err := loadWaterData(filePath, "AGUA", 3)
	if err != nil {
		fmt.Println("Error reading the water data" + err.Error())
	}

	for _, apar := range ret {
		err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, gpr["total_pres"], waterRead, waterData)
		if err != nil {
			fmt.Println(apar.number)
			fmt.Println(err)
		}
	}

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

func loadApartmentData(filePath, sheetName string) ([]Apartment, error) {
	// Open the spreadsheet
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

	cols := []string{}

	ret := []Apartment{}

out:
	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
			fmt.Println("Column information", cols)
		} else {
			ap := Apartment{}
		inside:
			for j, colCell := range row {
				switch strings.ToLower(cols[j]) {
				case "propietario":
					if len(colCell) == 0 {
						break out
					}
					ap.owner = colCell
				case "depa":
					ap.number = colCell
					// if err != nil {
					// 	fmt.Println("error en depa")
					// 	ap.number = 0.0
					// }
				case "área depa":
					ap.totalArea, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						fmt.Println("error en aread", ap.totalArea)
						ap.totalArea = 0.0
					}
				case "área est":
					ap.parkingArea, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.parkingArea = 0.0
					}
				case "total":
					ap.total, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.total = 0.0
					}
					break inside
				case "multa":
					ap.fine, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.fine = 0.0
					}
				case "cuota":
					ap.maintenance, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						ap.maintenance = 0.0
					}
				case "porcentaje":
					if len(colCell) == 0 {
						colCell = "--"
					}
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

	}
	return ret, nil
}

func loadWaterData(filePath, sheetName string, finalColumn int) (map[string]WaterMonthData, error) {
	// Open the spreadsheet
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
		fmt.Println(err)
		return nil, err
	}

	ret := make(map[string]WaterMonthData)

out:
	for i, row := range rows {
		if i == 0 || i == 1 || i == 2 {
			continue
		} else {
			var index string
			temp := WaterMonthData{}
			for j, colCell := range row {
				if j > finalColumn {
					break
				}
				switch j {
				case 0:
					if len(colCell) == 0 {
						break out
					}
					index = colCell
				case 1:
					temp.lastMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.lastMonth = 0.0
					}
				case 2:
					temp.currentMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.currentMonth = 0.0
					}
				case 3:
					temp.waterConsumedThisMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.waterConsumedThisMonth = 0.0
					}
				default:
					continue
				}
			}

			ret[index] = temp
		}
	}
	return ret, nil
}

func main() {
	// Limits in the spreadsheet

	filePath := "cuotas/GPR CUOTA OCTUBRE 2022.xlsx"

	sheetName := "Propietarios ordenados"

	ret, err := loadApartmentData(filePath, sheetName)
	if err != nil {
		panic(err)
	}

	var body bytes.Buffer

	email.GetTemplate("email/templates/maintenance.html", &body, "Octubre-2022")

	e := &email.EmailService{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "soporte-administrativo@elmolio.net",
	}
	err = e.SetNewDialer()
	if err != nil {
		panic(err)
	}

	err = e.Connect()
	if err != nil {
		panic(err)
	}

	for _, apar := range ret {
		allEmails := email.GetEmails()
		fmt.Println(allEmails[apar.number])
		err := e.SendReceipt(allEmails[apar.number], "Octubre-2022", "GPR-RECIBOS-OCTUBRE-2022/MANTENIMIENTO-OCTUBRE-2022_DPTO-"+apar.number+".pdf", &body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Email enviado exitosamente a " + apar.number)

	}

	e.Desconnect()
}
