package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"strconv"
	"strings"
	"time"

	"github.com/asloth/receipt-generator/apartment"
	"github.com/asloth/receipt-generator/building"
	"github.com/asloth/receipt-generator/email"
	"github.com/asloth/receipt-generator/fee"
	"github.com/asloth/receipt-generator/utils"
	"github.com/asloth/receipt-generator/water"
	"github.com/xuri/excelize/v2"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("PARA QUE SOY BUENO")
	fmt.Println("-----------------")
	fmt.Println("1. GENERAR RECIBOS")
	fmt.Println("2. ENVIAR RECIBOS POR CORREO")
	opti := "1"
	getData(reader, &opti)

	switch opti {
	case "1":
		generateRece(reader)
	case "2":
		sendEmails(reader)
	}

}

func generateRece(r *bufio.Reader) {
	reader := r
	fmt.Println("GENERAR RECIBOS")
	fmt.Println("---------------------")

	// Datos genrales que necesitamos del usuario
	fechaEmision := ""
	getReceiptData(reader, "fecha de emision (dd/mm/aa)", &fechaEmision, true)

	fechaVenc := ""
	getReceiptData(reader, "fecha de vencimiento (dd/mm/aa)", &fechaVenc, true)

	tipoCuota := "ORDINARIA"

	periodo := "AGOSTO-2022"
	getReceiptData(reader, "periodo", &periodo, false)

	fmt.Println("Ingrese el nombre del archivo excel, formato XLSX")
	name := "sheetName"
	getData(reader, &name)
	filePath := "cuotas/" + name + ".xlsx"

	fmt.Println("Ingrese el nombre de la hoja donde se encuentra la data de los departamentos")
	apartmentSheet := ""
	getData(reader, &apartmentSheet)

	fmt.Println("Ingrese el nombre de la hoja donde se encuentran los montos de cuotas")
	sheetName := "Propietarios ordenados"
	getData(reader, &sheetName)

	fmt.Println("ELIJA EL EDIFICIO DEL CUAL DESEA GENERAR RECIBOS")
	printBuilding()

	option := ""
	getData(reader, &option)
	waterRead := "BORRAR"
	var b building.Building
	switch option {
	case "2":
		b.GetBuildingData("belmonte")
	case "3":
		b.GetBuildingData("torrereal")
	case "4":
		b.GetBuildingData("mirador")
	case "5":
		b.GetBuildingData("nitoa")
	case "6":
		b.GetBuildingData("valera")
	case "7":
		b.GetBuildingData("golf")
	case "8":
		b.GetBuildingData("mora")
	case "9":
		b.GetBuildingData("alayza")
	case "10":
		b.GetBuildingData("sbs")
	case "11":
		b.GetBuildingData("montereal")
	case "12":
		b.GetBuildingData("tomasal")
	case "13":
		b.GetBuildingData("balcones")
	case "14":
		b.GetBuildingData("killa")
	case "15":
		b.GetBuildingData("gcc")
	case "16":
		b.GetBuildingData("elite")
	case "17":
		b.GetBuildingData("avila")
	case "18":
		b.GetBuildingData("huascar")
	case "19":
		b.GetBuildingData("rosapark")
	case "20":
		b.GetBuildingData("sanjose")
	case "21":
		b.GetBuildingData("rio")
		// TERMINA EL CASE
	}
	apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
	if err != nil {
		fmt.Println("Error reading aparment data" + err.Error())
	}
	fmt.Println("Directorio cargado")
	ret, err := fee.LoadFeeDetailData(filePath, sheetName)
	if err != nil {
		fmt.Println("Error reading fee data" + err.Error())
	}
	fmt.Println("Cuotas cargadas")
	waterData := make(map[string]water.WaterMonthData)
	waterGeneralData := &water.WaterByMonth{}
	addContometer := "n"
	fileExtension := "jpeg"
	if b.HaveWater {
		fmt.Println("Ingrese el nombre de la hoja donde se encuentran el agua POR DEPARTAMENTO")
		waterPath := "AGUA"
		getData(reader, &waterPath)
		waterData, err = loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
		
		fmt.Println("Desea agregar los contometros? (y/N)")
		getData(reader,&addContometer)
		
		fmt.Println("Ingrese la extension de las imagenes de los contometros (png/jpeg/jpg)")
		getData(reader,&fileExtension)
		
		fmt.Println("Ingrese el nombre de la hoja donde se encuentran los datos del recibo del agua")
		sheetNameWaterBuilding:=""
		getData(reader,&sheetNameWaterBuilding)
		waterGeneralData, err = utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding, 3)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
	} 
	for _, apar := range ret {
		err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData, &addContometer,&fileExtension)
		if err != nil {
			fmt.Println(apar.ApartmentNumber)
			fmt.Println(err)
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
				case 4:
					temp.CommonWater, err = strconv.ParseFloat(colCell, 64)
					if err != nil {
						temp.CommonWater = 0.0
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
func printBuilding() {
	fmt.Println("2. BELMONTE")
	fmt.Println("3. TORRE REAL")
	fmt.Println("4. MIRADOR")
	fmt.Println("5. NITOA")
	fmt.Println("6. VALERA")
	fmt.Println("7. GOLF PARK")
	fmt.Println("8. MORA")
	fmt.Println("9. ALAYZA")
	fmt.Println("10. SBS")
	fmt.Println("11. MONTE REAL")
	fmt.Println("12. TOMASAL")
	fmt.Println("13. BALCONES")
	fmt.Println("14. KILLA")
	fmt.Println("15. GRAN CENTRAL COLONIAL")
	fmt.Println("16. ELITE")
	fmt.Println("17. P. AVILA")
	fmt.Println("18. HUASCAR")
	fmt.Println("19. ROSA PARK")
	fmt.Println("20. CABALLERIZAS SAN JOSE")
	fmt.Println("21. RIO")
}
func getFilePath(reader *bufio.Reader) string {
	fmt.Println("Ingrese el nombre del archivo excel, formato XLSX")
	name := "sheetName"
	getData(reader, &name)
	filePath := "cuotas/" + name + ".xlsx"
	return filePath
}

func getSheetName(reader *bufio.Reader) string {
	fmt.Println("Ingrese el nombre de la hoja donde se encuentra la cuota")
	sheetName := "Propietarios ordenados"
	getData(reader, &sheetName)
	return sheetName
}
func getSheetDirectory(reader *bufio.Reader) string {
	fmt.Println("Ingrese el nombre de la hoja donde se encuentran los emails")
	emails := "DIRECTORIO"
	getData(reader, &emails)
	return emails
}
func getPeriodName(reader *bufio.Reader) string {
	fmt.Println("Ingrese el nombre del periodo al que pertenecen los recibos (Mes-A;o)")
	period := "Febrero-2023"
	getData(reader, &period)
	return period
}
func getApartmentNumber(reader *bufio.Reader) string {
	fmt.Println("Ingrese el numero del departamento que desea elegir")
	number := "Febrero-2023"
	getData(reader, &number)
	return number
}

func sendEmails(r *bufio.Reader) {
	reader := r
	fmt.Println("ENVIAR RECIBOS POR CORREO")
	fmt.Println("---------------------")

	filePath := getFilePath(reader)
	sheetName := getSheetName(reader)
	emails := getSheetDirectory(reader)
	period := getPeriodName(reader)

	fmt.Println("ELIJA EL EDIFICIO DEL CUAL DESEA ENVIAR LOS RECIBOS")
	printBuilding()
	option := ""
	getData(reader, &option)

	allemails, err := apartment.LoadAparmentData(filePath, emails)
	if err != nil {
		fmt.Println("Error reading apartment data" + err.Error())
	}

	var b building.Building

	switch option {
	// test := ret[len(ret)-19:] Utilizado para solo seleccionar a los ultimos 19 dptos
	case "2":
		b.GetBuildingData("belmonte")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}

		sendingEmail(ret, b, period, allemails)
	case "3":
		b.GetBuildingData("torrereal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "4":
		b.GetBuildingData("mirador")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "5":
		b.GetBuildingData("nitoa")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "6":
		b.GetBuildingData("valera")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "7":
		b.GetBuildingData("golf")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "8":
		b.GetBuildingData("mora")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "9":
		b.GetBuildingData("alayza")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "10":
		b.GetBuildingData("sbs")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "11":
		b.GetBuildingData("montereal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "12":
		b.GetBuildingData("tomasal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "13":
		b.GetBuildingData("balcones")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "14":
		b.GetBuildingData("killa")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "15":
		b.GetBuildingData("gcc")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "16":
		b.GetBuildingData("elite")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "17":
		b.GetBuildingData("avila")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "18":
		b.GetBuildingData("huascar")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "19":
		b.GetBuildingData("rosapark")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "20":
		b.GetBuildingData("sanjose")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	case "21":
		b.GetBuildingData("rio")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period, allemails)
	}

}

func getEmailConnection(period string, b *building.Building) (*email.EmailService, *bytes.Buffer) {
	var body bytes.Buffer
	email.GetTemplate("email/templates/maintenance.html", &body, period, b.Email)
	e := &email.EmailService{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "soporte-administrativo@elmolio.net",
	}
	err := e.SetNewDialer()
	if err != nil {
		panic(err)
	}

	err = e.Connect()
	if err != nil {
		panic(err)
	}
	return e, &body
}

func sendingEmail(ret []fee.FeeDetail, b building.Building, period string, allEmails []apartment.Apartment) {
	e, body := getEmailConnection(period, &b)

	for _, apar := range ret {
		email1 := apartment.GetItemByFieldValue(allEmails, apar.ApartmentNumber).FirstEmail
		email2 := apartment.GetItemByFieldValue(allEmails, apar.ApartmentNumber).SecondEmail

		fmt.Println("Enviando email a " + apar.ApartmentNumber + " con correo :" + email1)
		err := e.SendReceipt(email1, period, "output/"+b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apar.ApartmentNumber+".pdf", body)
		if len(email2) > 0 {
			fmt.Println("Enviando email a " + apar.ApartmentNumber + " con correo :" + email2)
			err = e.SendReceipt(email2, period, "output/"+b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apar.ApartmentNumber+".pdf", body)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Email enviado exitosamente a " + apar.ApartmentNumber)
	}
	e.Desconnect()
}

func sendEmailToAparment(period string, b building.Building, apartmentNumber string, allEmails []apartment.Apartment) {
	e, body := getEmailConnection(period, &b)
	email1 := apartment.GetItemByFieldValue(allEmails, apartmentNumber).FirstEmail
	email2 := apartment.GetItemByFieldValue(allEmails, apartmentNumber).SecondEmail
	fmt.Println("Enviando email a " + apartmentNumber + " con correo :" + email1)
	err := e.SendReceipt(email1, period, "output/"+b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apartmentNumber+".pdf", body)
	if len(email2) > 0 {
		fmt.Println("Enviando email a " + apartmentNumber + " con correo :" + email2)
		err = e.SendReceipt(email2, period, "output/"+b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apartmentNumber+".pdf", body)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email enviado exitosamente a " + apartmentNumber)

	e.Desconnect()
}
