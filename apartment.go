package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Apartment struct {
	number          int64
	owner           string
	totalArea       float64
	percentaje      float64
	total           float64
	maintenance     float64
	parking         string
	parkingArea     float64
	deposit         []string
	waterComsuption float64
}

func (ap *Apartment) GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, totalPresupuesto string) error {

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
	ReceiptHeader(&m, heightHeader)

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

	// Adding the header of DATOS DEL USUARIO section
	SubHeader(&m, colorMolio, "DATOS DEL PROPIETARIO/INQUILINO")

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
	ownerData := []string{ap.owner, strconv.Itoa(int(ap.number)), strconv.Itoa(int(ap.number)), string(ap.parking), string(ap.deposit[0]) + ", " + string(ap.deposit[1])}
	// Data for the second column of the receipt
	otherData := []string{dptoArea, parkingArea, participation + "%", fmt.Sprintf("%.2f", ap.waterComsuption), totalPresupuesto}

	// Reading the data and painting it into the receipt
	for i, v := range FirstColumn {
		DataOwner(&m, backgroundColor, rowHeight, contentSize, v, ownerData[i], SecondColumn[i], otherData[i])
	}

	// SECTION WATER DETAIL INFORMATION
	SubHeader(&m, colorMolio, "DETALLE DEL CONSUMO DE AGUA")
	// Defining the fields of the first column
	waterDetailsFirstColumn := []string{"PERIODO: ", "LECTURA ANTERIOR: ", "LECTURA ACTUAL: ", "CONSUMO: "}
	waterDetailsSecondColumn := []string{"CONSUMO REC: ", "S/. REC: ", "SOLES / M3	: ", " "}

	waterData := []string{"--", "--", "--", "--"}
	recData := []string{"--", "--", "--", " "}

	for i, fieldFirstColumn := range waterDetailsFirstColumn {
		DataOwner(&m, backgroundColor, rowHeight, contentSize, fieldFirstColumn, waterData[i], waterDetailsSecondColumn[i], recData[i])
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
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text("MANTENIMIENTO ",
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Left,
				})
		})
		m.Col(2, func() {
			m.Text(monto,
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text("AGUA ",
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Left,
				})
		})
		m.Col(2, func() {
			m.Text(fmt.Sprintf("S/. %.2f", ap.waterComsuption),
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})

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
	SubHeader(&m, colorMolio, "INFORMACION DE PAGO")
	PayInfo(&m, colorMolio)

	//FOOTER : AVISOS IMPORTANTES DE LA BOLETA
	SubHeader(&m, colorMolio, "AVISO IMPORTANTE")
	Footer(&m, backgroundColor, contentSize)

	// Create the directory to store the receipts
	if err := os.Mkdir("GPR-RECIBOS-"+periodo, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	// Create a custom name for the receipt
	fileName := "MANTENIMIENTO-" + periodo + "_DPTO-" + strconv.Itoa(int(ap.number)) + ".pdf"

	// Save the receipt into the directory
	err := m.OutputFileAndClose("GPR-RECIBOS-" + periodo + "/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func ReceiptHeader(pdf *pdf.Maroto, heightHeader float64) {
	m := *pdf
	var colWidth uint = 4

	m.Row(heightHeader, func() {
		m.Col(colWidth, func() {
			_ = m.FileImage("files/molio-logo.jpg", props.Rect{
				Center: true,
			})
		})
		m.Col(colWidth, func() {
			m.Text("RECIBO DE MANTENIMIENTO", props.Text{
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
				Top:   9,
				Color: color.Color{
					Red: 255,
				},
			})
			m.Text("CONDOMINIO GRAN PARQUE ROMA", props.Text{
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
				Top:   15,
			})
			m.Text("LEONARDO ARIETA 825 - CERCADO DE LIMA", props.Text{
				Size:  8,
				Style: consts.Bold,
				Align: consts.Center,
				Top:   20,
			})
		})
		m.Col(colWidth, func() {
			_ = m.FileImage("files/parque-roma-logo.jpg", props.Rect{
				Center: true,
			})
		})
	})
}

func DataOwner(pdf *pdf.Maroto, backgroundColor color.Color, rowHeight float64, contentSize float64, prop1, data1, prop2, data2 string) {
	m := *pdf
	m.SetBackgroundColor(backgroundColor)
	m.SetBorder(false)
	var column1 uint = 3
	var columnData uint = 4
	m.Row(9, func() {
		m.Col(column1, func() {
			m.Text(prop1, props.Text{
				Size:            9,
				Align:           consts.Left,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(data1, props.Text{
				Top: 1,
			})
		})
		m.Col(3, func() {
			m.Text(prop2, props.Text{
				Size:            9,
				Align:           consts.Left,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(2, func() {
			m.Text(data2, props.Text{
				Top: 1,
			})
		})
	})
}

func PayInfo(pdf *pdf.Maroto, colorMolio color.Color) {
	m := *pdf

	headers := []string{"BANCO", "CUENTA BANCARIA", "TITULAR DE CUENTA"}
	contents := [][]string{
		{"BCP", "3059864512041", "C. RECAUDADORA GRAN PARQUE ROMA"},
	}

	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      11.0,
			GridSizes: []uint{4, 4, 4},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      10.0,
			GridSizes: []uint{4, 4, 4},
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
}

func SubHeader(pdf *pdf.Maroto, colorMolio color.Color, subtitulo string) {
	m := *pdf
	m.SetBorder(true)
	m.SetBackgroundColor(colorMolio)
	m.Row(7, func() {
		m.Col(12, func() {
			m.Text(subtitulo,
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})
}

func Footer(pdf *pdf.Maroto, backgroundColor color.Color, contentSize float64) {
	m := *pdf
	m.SetBackgroundColor(backgroundColor)
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("1. Se deja por escrito que el incumplimiento del pago esta sujeto a mora.",
				props.Text{
					Size:  contentSize,
					Align: consts.Left,
				})
			m.Text("2. El Propietario autoriza el corte de suministro de agua por incumplimiento de pago.",
				props.Text{
					Size:  contentSize,
					Align: consts.Left,
					Top:   5,
				})
		})
	})
}
