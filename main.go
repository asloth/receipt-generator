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

func main(){
  reader := bufio.NewReader(os.Stdin)
	fmt.Println("PARA QUE SOY BUENO")
	fmt.Println("-----------------")
	fmt.Println("1. GENERAR RECIBOS")
  fmt.Println("2. ENVIAR RECIBOS POR CORREO")
  opti := "1"
	getData(reader, &opti)

  switch opti{
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

	tipoCuota := "ORDINARIO"
	getReceiptData(reader, "tipo de cuota", &tipoCuota, false)

	periodo := "AGOSTO-2022"
	getReceiptData(reader, "periodo", &periodo, false)

	waterRead := ""
	getReceiptData(reader, "fecha de lectura del agua (dd/mm/aa)", &waterRead, true)

	fmt.Println("Ingrese el nombre del archivo excel, formato XLSX")
	name := "sheetName"
	getData(reader, &name)
	filePath := "cuotas/" + name + ".xlsx"

	fmt.Println("Ingrese el nombre de la hoja donde se encuentra la data de los departamentos")
	apartmentSheet := ""
	getData(reader, &apartmentSheet)

	fmt.Println("Ingrese el nombre de la hoja donde se encuentran el agua POR DEPARTAMENTO")
	waterPath := "AGUA"
	getData(reader, &waterPath)

	fmt.Println("Ingrese el nombre de la hoja donde se encuentran los montos de cuotas")
	sheetName := "Propietarios ordenados"
	getData(reader, &sheetName)

  fmt.Println("Ingrese el nombre donde estan los datos de agua del recibo")
  sheetNameWaterBuilding := ""
  getData(reader, &sheetNameWaterBuilding)

	fmt.Println("ELIJA EL EDIFICIO DEL CUAL DESEA GENERAR RECIBOS")
	fmt.Println("2. BELMONTE")
	fmt.Println("3. TORRE REAL")
	fmt.Println("4. MIRADOR")
	fmt.Println("5. NITOA")
	fmt.Println("6. VALERA")
	fmt.Println("7. GOLF PARK")
	fmt.Println("8. MORA")
	fmt.Println("9. ALAYZA")
	fmt.Println("10. SAN BORJA SUR")
	fmt.Println("11. MONTE REAL")
	fmt.Println("12. TOMASAL")
	fmt.Println("13. BALCONES")
	fmt.Println("14. KILLA")
	fmt.Println("15. GRAN CENTRAL COLONIAL")
	option := ""
	getData(reader, &option)

	var b building.Building
	switch option {
	case "2":
		b.GetBuildingData("belmonte")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		} 
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "3":
		b.GetBuildingData("torrereal")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		} 
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "4":
		b.GetBuildingData("mirador")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "5":
		b.GetBuildingData("nitoa")
		apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
		if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
		waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
		fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "6":
		b.GetBuildingData("valera")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "7":
		b.GetBuildingData("golf")
		apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
		if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
		waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
		fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "8":
		b.GetBuildingData("mora")
		apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
		if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
		waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
		fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "9":
		b.GetBuildingData("alayza")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "10":
		b.GetBuildingData("sbs")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "11":
		b.GetBuildingData("montereal")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    if err != nil { 
			fmt.Println("Error reading aparment data" + err.Error())
		}
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)

		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
   
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		} 
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "12":
		b.GetBuildingData("tomasal")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    fmt.Println("soy appdata" , apData)
   	if err != nil { 
			fmt.Println("Error reading apartment data" + err.Error())
		}
    fmt.Println(apData)
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
    fmt.Println("soy ret" , ret)

		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
    fmt.Println("soy waterData" , waterData)

		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}
    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}

		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
  case "13":
		b.GetBuildingData("balcones")
    apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
    fmt.Println("soy appdata" , apData)
   	if err != nil { 
			fmt.Println("Error reading apartment data" + err.Error())
		}
    fmt.Println(apData)
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
    fmt.Println("soy ret" , ret)

		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
    fmt.Println("soy waterData" , waterData)
    if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}

    waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
    fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "14":
		b.GetBuildingData("killa")
		apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
		fmt.Println("soy appdata" , apData)
		if err != nil { 
			fmt.Println("Error reading apartment data" + err.Error())
		}
		fmt.Println(apData)
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		fmt.Println("soy ret" , ret)

		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		fmt.Println("soy waterData" , waterData)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}

		waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
		fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}
	case "15":
		b.GetBuildingData("gcc")
		apData, err := apartment.LoadAparmentData(filePath, apartmentSheet)
		fmt.Println("soy appdata" , apData)
		if err != nil { 
			fmt.Println("Error reading apartment data" + err.Error())
		}
		fmt.Println(apData)
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		fmt.Println("soy ret" , ret)

		if err != nil {
			fmt.Println("Error reading fee data" + err.Error())
		}
		waterData, err := loadWaterData(filePath, waterPath, 4)
		fmt.Println("soy waterData" , waterData)
		if err != nil {
			fmt.Println("Error reading the water data" + err.Error())
		}

		waterGeneralData,err := utils.LoadWaterBuilding(filePath, sheetNameWaterBuilding,3)
		fmt.Println("soy waterGeneralData",waterGeneralData)

		if err != nil {
			fmt.Println("Error reading the water general data" + err.Error())
		}
		for _, apar := range ret {
			err := apar.GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterRead, waterData, &b, &apData, *waterGeneralData)
			if err != nil {
				fmt.Println(apar.ApartmentNumber)
				fmt.Println(err)
			}
		}

    // TERMINA EL CASE
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
					fmt.Println("i am consumption: ",colCell)
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

func sendEmails( r *bufio.Reader) {
	reader := r
	fmt.Println("ENVIAR RECIBOS POR CORREO")
	fmt.Println("---------------------")

	fmt.Println("Ingrese el nombre del archivo excel, formato XLSX")
	name := "sheetName"
	getData(reader, &name)
	filePath := "cuotas/" + name + ".xlsx"

	fmt.Println("Ingrese el nombre de la hoja donde se encuentran los propietarios ordenados")
	sheetName := "Propietarios ordenados"
	getData(reader, &sheetName)

	fmt.Println("Ingrese el nombre del periodo al que pertenecen los recibos (Mes-A;o)")
	period := "Febrero-2023"
	getData(reader, &period)

	fmt.Println("ELIJA EL EDIFICIO DEL CUAL DESEA ENVIAR LOS RECIBOS")
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

	option := ""
	getData(reader, &option)

	var b building.Building

	switch option {
		// test := ret[len(ret)-19:] Utilizado para solo seleccionar a los ultimos 19 dptos
	case "2":
		b.GetBuildingData("belmonte")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "3":
		b.GetBuildingData("torrereal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "4":
		b.GetBuildingData("mirador")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "5":
		b.GetBuildingData("nitoa")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "6":
		b.GetBuildingData("valera")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "7":
		b.GetBuildingData("golf")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "8":
		b.GetBuildingData("mora")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "9":
		b.GetBuildingData("alayza")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "10":
		b.GetBuildingData("sbs")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "11":
		b.GetBuildingData("montereal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "12":
		b.GetBuildingData("tomasal")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "13":
		b.GetBuildingData("balcones")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "14":
		b.GetBuildingData("killa")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)
	case "15":
		b.GetBuildingData("gcc")
		ret, err := fee.LoadFeeDetailData(filePath, sheetName)
		if err != nil {
			panic(err)
		}
		sendingEmail(ret, b, period)

	}

}

func sendingEmail(ret []fee.FeeDetail, b building.Building, period string) {
	var body bytes.Buffer
	//CAMBIAR NOMBRE CUOTA
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

	for _, apar := range ret {
		allEmails := *email.GetEmails(b.Nickname)
		fmt.Println("Enviando email a " + apar.ApartmentNumber + " con correo :" + allEmails[apar.ApartmentNumber][0])
		err := e.SendReceipt(allEmails[apar.ApartmentNumber][0], period, "output/"+ b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apar.ApartmentNumber+".pdf", &body)
		if len(allEmails[apar.ApartmentNumber][1]) > 0 {
			fmt.Println("Enviando email a " + apar.ApartmentNumber + " con correo :" + allEmails[apar.ApartmentNumber][1])
			err = e.SendReceipt(allEmails[apar.ApartmentNumber][1], period, "output/"+b.Nickname+"-RECIBOS-"+strings.ToUpper(period)+"/MANTENIMIENTO-"+strings.ToUpper(period)+"_DPTO-"+apar.ApartmentNumber+".pdf", &body)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Email enviado exitosamente a " + apar.ApartmentNumber)

	}

	e.Desconnect()
}

