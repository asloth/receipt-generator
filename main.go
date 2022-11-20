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

	"github.com/asloth/receipt-generator/building"
	"github.com/asloth/receipt-generator/email"
	"github.com/asloth/receipt-generator/fee"
	"github.com/asloth/receipt-generator/water"
	"github.com/xuri/excelize/v2"
)

func mainreceipts() {

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

	fmt.Println("ELIJA EL EDIFICIO DEL CUAL DESEA GENERAR RECIBOS")
	fmt.Println("1. GRAN PARQUE ROMA")
	fmt.Println("2. BELMONTE")

	option := ""
	getData(reader, &option)

	filePath := "cuotas/BELMONTE CUOTA NOVIEMBRE 2022.xlsx"

	waterPath := "AGUA"

	sheetName := "Propietarios ordenados"

	var b building.Building
	switch option {
	case "1":
		b.GetBuildingData("gpr")
		b.Budget = totalPresupuesto

		ret, err := loadApartmentData(filePath, sheetName)

		if err != nil {
			fmt.Println("Error reading apartment data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 3)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}

		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b)
			if err != nil {
				fmt.Println(apar.number)
				fmt.Println(err)
			}
		}
	case "2":
		b.GetBuildingData("belmonte")
		b.Budget = totalPresupuesto

		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 3)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}

		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	}

}

func getData(r *bufio.Reader, data *string) {
	reader := *r
	for {

		// Reading the user input
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		stringText := strings.Replace(text, "\n", "", -1)

		*data = stringText

		break

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

func loadWaterData(filePath, sheetName string, finalColumn int) (map[string]water.WaterMonthData, error) {
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

	ret := make(map[string]water.WaterMonthData)

out:
	for i, row := range rows {
		if i == 0 || i == 1 || i == 2 {
			continue
		} else {
			var index string
			temp := water.WaterMonthData{}
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
					temp.LastMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.LastMonth = 0.0
					}
				case 2:
					temp.CurrentMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.CurrentMonth = 0.0
					}
				case 3:
					temp.WaterConsumedThisMonth, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.WaterConsumedThisMonth = 0.0
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

	var b building.Building
	b.GetBuildingData("belmonte")

	filePath := "cuotas/BELMONTE CUOTA NOVIEMBRE 2022.xlsx"

	sheetName := "Propietarios ordenados"

	ret, err := fee.LoadFeeDetailData(filePath, sheetName)
	if err != nil {
		panic(err)
	}

	var body bytes.Buffer

	email.GetTemplate("email/templates/maintenance.html", &body, "Noviembre-2022", b.Email)

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
		allEmails := *email.GetEmails(b.Nickname)
		fmt.Println(allEmails[apar.ApartmentNumber])
		err := e.SendReceipt(allEmails[apar.ApartmentNumber], "Noviembre-2022", b.Nickname+"-RECIBOS-NOVIEMBRE-2022/MANTENIMIENTO-NOVIEMBRE-2022_DPTO-"+apar.ApartmentNumber+".pdf", &body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Email enviado exitosamente a " + apar.ApartmentNumber)

	}

	e.Desconnect()
}
