package fee

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/asloth/receipt-generator/apartment"
	"github.com/asloth/receipt-generator/building"
	"github.com/asloth/receipt-generator/receipt"
	"github.com/asloth/receipt-generator/water"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/xuri/excelize/v2"
)

type FeeDetail struct {
	ApartmentNumber string
	Amounts         map[string]float64
}

func LoadFeeDetailData(filePath, sheetName string) ([]FeeDetail, error) {
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

	ret := []FeeDetail{}

	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
		} else {
			ap := FeeDetail{}
			ap2 := make(map[string]float64)
		inside:
			for j, colCell := range row {
				colCell = strings.TrimSpace(colCell)               //el valor de la celda
				col := strings.TrimSpace(strings.ToLower(cols[j])) //el nombre de la columna
				if j == 0 {
					ap.ApartmentNumber = colCell
					continue inside
				}
				ap2[col], err = strconv.ParseFloat(colCell, 64)
			}
			ap.Amounts = ap2
			ret = append(ret, ap)
		}

	}
	return ret, nil
}

func findApartmentByID(id string, myAp []apartment.Apartment) *apartment.Apartment {
	for i := range myAp {
		if myAp[i].Number == id {
			return &myAp[i]
		}
	}
	return nil // Return nil if the struct with the given ID is not found
}

func (ap *FeeDetail) GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterDate string, wData map[string]water.WaterMonthData, b *building.Building, apData *[]apartment.Apartment, wGeneralData water.WaterByMonth, cont *string, fileEx *string) error {
	apList := *apData

	myAp := findApartmentByID(ap.ApartmentNumber, apList)

	if myAp == nil {
		return errors.New("Numero de departamento no existe")
	}
	buildng := *b
	var heightHeader float64 = 30
	var contentSize float64 = 10
	var rowHeight float64 = 7
	colorMolio := &props.Color{
		Red:   148,
		Green: 235, //bajar un poco hasta 200
		Blue:  66,
	}
	cfg := config.NewBuilder().
	        WithConcurrentMode(7).
		Build()
	m := maroto.New(cfg)
	// Header
	addCont := parseYesNo(*cont)
	pathCont := "contometer/"+buildng.Nickname+"/"+periodo+"/"+ap.ApartmentNumber+"."+*fileEx
	if addCont {
		if !fileExists(pathCont) {
			pathCont = b.Picture
		}
	}
	receipt.ReceiptHeader(&m, heightHeader, &buildng, &pathCont)
	
	m.AddRow(2) 
	m.AddRow(5,line.NewCol(12)) 

	// tabla inicial
	colStyleHeader := &props.Cell{
		BackgroundColor: colorMolio,
	        BorderType:      border.Full,
		BorderColor:     &props.BlackColor,
		BorderThickness: 0.3,
	}

	colStyleCenterContent := &props.Cell{
	        BorderType:      border.Full,
		BorderColor:     &props.BlackColor,
		BorderThickness: 0.3,
	}
	headerText := props.Text{
		Family:    fontfamily.Arial,
		Style:     fontstyle.Bold,
		Size:      11.0,
		Top: 0.5,
		Align: align.Center,
		Color: &props.BlackColor,
	}

	headers := []core.Col{
		text.NewCol(3,"TIPO CUOTA",headerText).WithStyle(colStyleHeader),
		text.NewCol(3,"F. EMISION",headerText).WithStyle(colStyleHeader),
		text.NewCol(3,"F. VCTO.",headerText).WithStyle(colStyleHeader),
		text.NewCol(3,"PERIODO",headerText).WithStyle(colStyleHeader),
	}

	contentCenterCell := props.Text{
		Family:    fontfamily.Courier,
		Style:     fontstyle.Normal,
		Top: 1,
		Size:      10.0,
		Align: align.Center,
		Color: &props.BlackColor,
	}
	contents := []core.Col{
		text.NewCol(3,tipoCuota,contentCenterCell).WithStyle(colStyleCenterContent),
		text.NewCol(3,fechaEmision,contentCenterCell).WithStyle(colStyleCenterContent),
		text.NewCol(3,fechaVenc,contentCenterCell).WithStyle(colStyleCenterContent),
		text.NewCol(3,periodo,contentCenterCell).WithStyle(colStyleCenterContent),
	}
	m.AddRow(7,headers...)	
	m.AddRow(9,contents...)	
	// SECCION DATOS DEL DPTO
	receipt.SubHeader(&m, "DATOS DEL DEPARTAMENTO", colStyleHeader)

	printAparmentData(&m, colStyleCenterContent, contentSize, myAp)
	// SECTION DATOS DEL USUARIO
	receipt.SubHeader(&m, "DETALLE DE LA CUOTA",colStyleHeader)

	Detail(&m, contentSize, rowHeight, ap, myAp)

	// SECTION WATER DETAIL INFORMATION
	if buildng.HaveWater {
		receipt.SubHeader(&m,"CONSUMOS INDIVIDUALES", colStyleHeader)
		// Defining the fields of the first column
		waterDetailsFirstColumn := []string{"CONSUMO COMUN: ", "LECTURA ANTERIOR: ", "LECTURA ACTUAL: ", "CONSUMO: "}
		waterDetailsSecondColumn := []string{"CONSUMO REC: ", "S/. REC: ", "COSTO UNITARIO: ", ""}
		waterData := []string{fmt.Sprintf("S/. %.2f", wData[ap.ApartmentNumber].CommonWater), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].LastMonth), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].CurrentMonth), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].WaterConsumedThisMonth)}

		// Get water data from this month
		recData := []string{wGeneralData.Consumo_rec, wGeneralData.Rec_soles, wGeneralData.Soles_m3, ""}

		for i, fieldFirstColumn := range waterDetailsFirstColumn {
			receipt.DataOwner(&m, rowHeight, contentSize, fieldFirstColumn, waterData[i], waterDetailsSecondColumn[i], recData[i])
		}
	}

	//IMPORTES FACTURADOS SECTION TABLE
	monto := fmt.Sprintf("%.2f", ap.Amounts["cuota"])
	resumenTextProps := props.Text{
				Size:  12,
				Style: fontstyle.Bold,
				Top: 0.5,
				Align: align.Center,
			}
	m.AddRow(7, 
		text.NewCol(10,"IMPORTES FACTURADOS", resumenTextProps).WithStyle(colStyleHeader),
		text.NewCol(2,"IMPORTE",resumenTextProps).WithStyle(colStyleHeader),
	)

	receipt.Resumen(&m, colStyleCenterContent, contentSize, "MANTENIMIENTO ", monto)

	m.AddRow(7, 
		text.NewCol(10,"TOTAL A PAGAR S/.", resumenTextProps).WithStyle(colStyleHeader),
		text.NewCol(2,monto,resumenTextProps).WithStyle(colStyleHeader),
	)
	m.AddRow(7)

	// PAY INFORMACION
	receipt.SubHeader(&m, "INFORMACION DE PAGO", colStyleHeader)
	receipt.PayInfo(&m, colStyleHeader,colStyleCenterContent, &buildng)


	// Create the directory to store the receipts
	if err := os.Mkdir("output/"+buildng.Nickname+"-RECIBOS-"+periodo, os.ModePerm); err != nil {

	}

	// Create a custom name for the receipt
	fileName := "MANTENIMIENTO-" + periodo + "_DPTO-" + ap.ApartmentNumber + ".pdf"

	// Save the receipt into the directory
	document, err := m.Generate()
	if err != nil {
		fmt.Println(err)
	}
	faildoc := document.Save("output/" + buildng.Nickname + "-RECIBOS-" + periodo + "/" + fileName)

	if err != nil {
		fmt.Println(faildoc)
	}

	return nil
}

func Detail(pdf *core.Maroto, contentSize, rowHeight float64, ap *FeeDetail, myApartment *apartment.Apartment) {
	m := *pdf
	totalItems := len(ap.Amounts) - 1
	var ownerData []string
	var otherData []string
	var itemsByColumn int = totalItems / 2
	var FirstColumn []string  // Defining the fields for the first column of the receipt
	var SecondColumn []string // Defining the fields for the second column of the receipt

	if totalItems%2 != 0 {
		itemsByColumn++ // La cantidad de elementos que iran por columna
	}
	var j int = 0
	for key, value := range ap.Amounts {
		if key == "cuota" {
			continue
		}

		if j < itemsByColumn {
			FirstColumn = append(FirstColumn, key)
			ownerData = append(ownerData, fmt.Sprintf("%.2f", value))
			j++
			continue
		}
		SecondColumn = append(SecondColumn, key)
		otherData = append(otherData, fmt.Sprintf("%.2f", value))
	}

	if len(FirstColumn) != len(SecondColumn) {
		SecondColumn = append(SecondColumn, " ")
		otherData = append(otherData, " ")
	}

	// Reading the data and painting it into the receipt
	for i, v := range FirstColumn {
		receipt.DataOwner(&m, rowHeight, contentSize, v, ownerData[i], SecondColumn[i], otherData[i])
	}
}

func parseYesNo(input string) bool {
	normalizedInput := strings.TrimSpace(input) // Trim surrounding spaces
	normalizedInput = strings.ToUpper(normalizedInput) // Convert to uppercase

	if normalizedInput == "Y" {
		return true
	} else if normalizedInput == "N" {
		return false
	}
	
	// Handle unexpected input (optional)
	fmt.Printf("Invalid input: %s. Please provide 'y' or 'n'.\n", input)
	return false
}

func printAparmentData(pdf *core.Maroto, colSyleCenterContent *props.Cell, contentSize float64, ap *apartment.Apartment) {
	// Get the type of the struct
	m := *pdf
	structType := reflect.TypeOf(*ap)
	fieldName := []string{"N. DPTO: ", "PROPIETARIO: ", "ESTACIONAMIENTO: ", "DEPOSITO: ", "PARTICIPACION: "}

	// Loop over the struct fields
	for i := 0; i < structType.NumField(); i++ {
		if i > 4 {
			continue
		}
		fieldValue := reflect.ValueOf(*ap).Field(i)
		if len(fieldValue.String()) == 0 {
			continue
		}
		receipt.ApartmentData(&m, colSyleCenterContent, contentSize, fieldName[i], fieldValue.String())
	}
}

// Function to check if a file exists at the given path
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
