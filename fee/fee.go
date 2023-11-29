package fee

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/asloth/receipt-generator/building"
	"github.com/asloth/receipt-generator/receipt"
	"github.com/asloth/receipt-generator/water"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/xuri/excelize/v2"
)

type FeeDetail struct {
	ApartmentNumber      string
  Amounts map[string]float64
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
			fmt.Println("Column information", cols)
		} else {
      ap := FeeDetail{}
      ap2 := make(map[string]float64)
      inside:
			for j, colCell := range row {
				colCell = strings.TrimSpace(colCell) //el valor de la celda 
        col := strings.TrimSpace(strings.ToLower(cols[j])) //el nombre de la columna
        if j == 0 {
          ap.ApartmentNumber = colCell
          continue inside
        }
        ap2[col], err = strconv.ParseFloat(colCell, 64)
        fmt.Println("soy ap2", ap2)
			}
      ap.Amounts = ap2
		}

	}
	return ret, nil
}

func (ap *FeeDetail) GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterDate string, wData map[string]water.WaterMonthData, b *building.Building) error {
	buildng := *b
	var heightHeader float64 = 30
	var contentSize float64 = 10
	var rowHeight float64 = 7
	colorMolio := color.Color{
		Red:   148,
		Green: 235,
		Blue:  66,
	}
	backgroundColor := color.NewWhite()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	// Header
	receipt.ReceiptHeader(&m, heightHeader, &buildng)

	// tabla inicial
	headers := []string{"TIPO CUOTA", "F. EMISION", "F. VCTO.", "PERIODO"}
	contents := [][]string{
		{tipoCuota, fechaEmision, fechaVenc, periodo},
	}
	m.Line(10)
	m.SetBorder(true)
	m.SetBackgroundColor(colorMolio)

	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      11.0,
			GridSizes: []uint{3, 3, 3, 3},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      10.0,
			GridSizes: []uint{3, 3, 3, 3},
		},
		Align:                  consts.Center,
		HeaderContentSpace:     0.01,
		VerticalContentPadding: 4.0,
		AlternatedBackground: &color.Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		},
	})

	// SECTION DATOS DEL USUARIO
	receipt.SubHeader(&m, colorMolio, "DETALLE DEL CONSUMO DE LA CUOTA")

	Detail(&m, backgroundColor, contentSize, rowHeight, ap, &buildng)

	// SECTION WATER DETAIL INFORMATION
	if buildng.HaveWater {
		receipt.SubHeader(&m, colorMolio, "DETALLE DEL CONSUMO DE AGUA")
		// Defining the fields of the first column
		waterDetailsFirstColumn := []string{"AGUA COMUN: ", "LECTURA ANTERIOR (m3): ", "LECTURA ACTUAL (m3): ", "CONSUMO (m3): "}
		waterDetailsSecondColumn := []string{"CONSUMO REC: ", "S/. REC: ", "SOLES / M3: ", "FECHA DE LECTURA: "}

		waterData := []string{fmt.Sprintf("S/. %.2f", wData[ap.ApartmentNumber].CommonWater), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].LastMonth), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].CurrentMonth), fmt.Sprintf("%.2f", wData[ap.ApartmentNumber].WaterConsumedThisMonth)}

		// Get water data from this month
		monthWaterData := water.GetWaterDataByBuilding(b.Nickname)
		recData := []string{fmt.Sprintf("%.2f", monthWaterData.Consumo_rec), fmt.Sprintf("%.2f", monthWaterData.Rec_soles), fmt.Sprintf("%.2f", monthWaterData.Soles_m3), waterDate}

		for i, fieldFirstColumn := range waterDetailsFirstColumn {
			receipt.DataOwner(&m, backgroundColor, rowHeight, contentSize, fieldFirstColumn, waterData[i], waterDetailsSecondColumn[i], recData[i])
		}
	}

	//IMPORTES FACTURADOS SECTION TABLE
	monto := fmt.Sprintf("S/. %.2f", ap.Amounts["CUOTA"])
	m.SetBackgroundColor(colorMolio)
	m.SetBorder(true)
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text("IMPORTES FACTURADOS",
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
		m.Col(2, func() {
			m.Text("IMPORTE",
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})

	m.SetBackgroundColor(backgroundColor)
	receipt.Resumen(&m, backgroundColor, contentSize, "MANTENIMIENTO ", monto)

	m.SetBackgroundColor(colorMolio)
	m.SetBorder(true)
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text("TOTAL A PAGAR S/.",
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
		m.Col(2, func() {
			m.Text(monto,
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})
	m.SetBackgroundColor(backgroundColor)
	m.Row(7, func() {})

	// PAY INFORMACION
	receipt.SubHeader(&m, colorMolio, "INFORMACION DE PAGO")
	receipt.PayInfo(&m, colorMolio, &buildng)

	//FOOTER : AVISOS IMPORTANTES DE LA BOLETA
	//receipt.SubHeader(&m, colorMolio, "AVISO IMPORTANTE")
	//receipt.Footer(&m, backgroundColor, contentSize)

	// Create the directory to store the receipts
	if err := os.Mkdir(buildng.Nickname+"-RECIBOS-"+periodo, os.ModePerm); err != nil {

	}

	// Create a custom name for the receipt
	fileName := "MANTENIMIENTO-" + periodo + "_DPTO-" + ap.ApartmentNumber + ".pdf"

	// Save the receipt into the directory
	err := m.OutputFileAndClose(buildng.Nickname + "-RECIBOS-" + periodo + "/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func Detail(pdf *pdf.Maroto, backgroundColor color.Color, contentSize, rowHeight float64, ap *FeeDetail, buildng *building.Building) {
	m := *pdf
	var ownerData []string
	var otherData []string

	// Defining the fields for the first column of the receipt
	FirstColumn := buildng.FirstColumn

	// Defining the fields for the second column of the receipt
	SecondColumn := buildng.SecondColumn

	switch strings.ToLower(buildng.Nickname) {
	
	}

	// Reading the data and painting it into the receipt
	for i, v := range FirstColumn {
		receipt.DataOwner(&m, backgroundColor, rowHeight, contentSize, v, ownerData[i], SecondColumn[i], otherData[i])
	}
}
