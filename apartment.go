package main

import (
	"fmt"
	"os"

	"github.com/asloth/receipt-generator/building"
	"github.com/asloth/receipt-generator/receipt"
	"github.com/asloth/receipt-generator/water"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Apartment struct {
	number          string
	owner           string
	totalArea       float64
	percentaje      float64
	total           float64
	maintenance_ext float64
	maintenance     float64
	parking         string
	parkingArea     float64
	deposit         []string
	waterComsuption float64
	fine            float64
}

func (ap *Apartment) GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, waterDate string, wData map[string]water.WaterMonthData, b *building.Building) error {
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
	receipt.SubHeader(&m, colorMolio, "DATOS DEL PROPIETARIO/INQUILINO")
	UserDetail(&m, backgroundColor, contentSize, rowHeight, ap, &buildng)

	// SECTION WATER DETAIL INFORMATION
	receipt.SubHeader(&m, colorMolio, "DETALLE DEL CONSUMO DE AGUA")
	// Defining the fields of the first column
	waterDetailsFirstColumn := []string{"PERIODO: ", "LECTURA ANTERIOR (m3): ", "LECTURA ACTUAL (m3): ", "CONSUMO (m3): "}
	waterDetailsSecondColumn := []string{"CONSUMO REC: ", "S/. REC: ", "SOLES / M3: ", "FECHA DE LECTURA: "}

	waterData := []string{periodo, fmt.Sprintf("%.2f", wData[ap.number].LastMonth), fmt.Sprintf("%.2f", wData[ap.number].CurrentMonth), fmt.Sprintf("%.2f", wData[ap.number].WaterConsumedThisMonth)}

	// Get water data from this month
	monthWaterData := water.GetWaterDataByBuilding(b.Nickname)
	recData := []string{fmt.Sprintf("%.2f", monthWaterData.Consumo_rec), fmt.Sprintf("%.2f", monthWaterData.Rec_soles), fmt.Sprintf("%.2f", monthWaterData.Soles_m3), waterDate}

	for i, fieldFirstColumn := range waterDetailsFirstColumn {
		receipt.DataOwner(&m, backgroundColor, rowHeight, contentSize, fieldFirstColumn, waterData[i], waterDetailsSecondColumn[i], recData[i])
	}

	//IMPORTES FACTURADOS SECTION TABLE
	monto := fmt.Sprintf("S/. %.2f", ap.maintenance)

	m.SetBackgroundColor(colorMolio)
	m.SetBorder(true)
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text("DETALLE DE LOS IMPORTES FACTURADOS",
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
	receipt.Resumen(&m, backgroundColor, contentSize, "AGUA ", fmt.Sprintf("S/. %.2f", ap.waterComsuption))
	receipt.Resumen(&m, backgroundColor, contentSize, "MULTA ", fmt.Sprintf("S/. %.2f", ap.fine))
	receipt.Resumen(&m, backgroundColor, contentSize, "CUOTA EXTRAORDINARIA ", fmt.Sprintf("S/. %.2f", ap.maintenance_ext))

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
			m.Text(fmt.Sprintf("S/. %.2f", ap.total),
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
	receipt.SubHeader(&m, colorMolio, "AVISO IMPORTANTE")
	receipt.Footer(&m, backgroundColor, contentSize)

	// Create the directory to store the receipts
	if err := os.Mkdir(buildng.Nickname+"-RECIBOS-"+periodo, os.ModePerm); err != nil {

	}

	// Create a custom name for the receipt
	fileName := "MANTENIMIENTO-" + periodo + "_DPTO-" + ap.number + ".pdf"

	// Save the receipt into the directory
	err := m.OutputFileAndClose(buildng.Nickname + "-RECIBOS-" + periodo + "/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func UserDetail(pdf *pdf.Maroto, backgroundColor color.Color, contentSize, rowHeight float64, ap *Apartment, buildng *building.Building) {
	m := *pdf

	// Defining the fields for the first column of the receipt
	FirstColumn := []string{"NOMBRE: ", "DEPARTAMENTO: ", "CODIGO BANCO: ", "ESTACIONAMIENTO: ", "DEPOSITO: "}

	// Defining the fields for the second column of the receipt
	SecondColumn := []string{
		"AREA DEPARTAMENTO: ",
		"AREA ESTACIONAMIENTO: ",
		"% PARTICIPACION: ",
		"CONSUMO AGUA (S/.): ",
		"TOTAL PRESUPUESTO: "}

	// Parsing data from float to string with 2 decimals to show in the receipt
	dptoArea := fmt.Sprintf("%.2f m2", ap.totalArea)
	parkingArea := fmt.Sprintf("%.2f m2", ap.parkingArea)
	participation := fmt.Sprintf("%f", ap.percentaje)

	// Data for the first column of the receipt
	ownerData := []string{ap.owner, ap.number, ap.number, string(ap.parking), string(ap.deposit[0]) + ", " + string(ap.deposit[1])}
	// Data for the second column of the receipt
	otherData := []string{dptoArea, parkingArea, participation + "%", fmt.Sprintf("%.2f", ap.waterComsuption), buildng.Budget}

	// Reading the data and painting it into the receipt
	for i, v := range FirstColumn {
		receipt.DataOwner(&m, backgroundColor, rowHeight, contentSize, v, ownerData[i], SecondColumn[i], otherData[i])
	}
}
