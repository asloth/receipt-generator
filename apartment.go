package main

import (
	"fmt"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Apartment struct {
	number     int64
	owner      string
	totalArea  float64
	percentaje float64
	amount     float64
	parking    string
}

func (ap *Apartment) GenerateReceipt(tipoCuota, fechaEmision, fechaVenc, periodo, totalPresupuesto string) {

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

	// tabla central datos del usuario
	SubHeader(&m, colorMolio, "DATOS DEL PROPIETARIO/INQUILINO")

	atributes := []string{"NOMBRE: ", "DEPARTAMENTO: ", "CODIGO BANCO: ", "AREA DEPARTAMENTO: ", "ESTACIONAMIENTO: ", "% PARTICIPACION: ", "TOTAL PRESUPUESTO: "}

	dptoArea := fmt.Sprintf("%.2f m2", ap.totalArea)
	participation := fmt.Sprintf("%f", ap.percentaje)

	monto := fmt.Sprintf("S/. %.2f", ap.amount)

	ownerData := []string{ap.owner, strconv.Itoa(int(ap.number)), strconv.Itoa(int(ap.number)), dptoArea, string(ap.parking), participation + "%", totalPresupuesto}

	m.Row(20, func() {
		m.Col()
		m.Col()
	})
	for i, v := range atributes {
		DataOwner(&m, backgroundColor, rowHeight, contentSize, v, ownerData[i])
	}

	//TERCERA TABLA de importes fcturados
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
			m.Text("MANTENIMIENTO ( "+participation+" % ) x ( S/ "+totalPresupuesto+" )",
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

	SubHeader(&m, colorMolio, "INFORMACION DE PAGO")
	PayInfo(&m, colorMolio)

	//FOOTER : AVISOS IMPORTANTES DE LA BOLETA
	SubHeader(&m, colorMolio, "AVISO IMPORTANTE")
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
	fileName := "Mantenimiento-" + periodo + "_DPTO-01-" + string(rune(ap.number)) + ".pdf"
	m.OutputFileAndClose(fileName)
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

func DataOwner(pdf *pdf.Maroto, backgroundColor color.Color, rowHeight float64, contentSize float64, prop, data string) {
	m := *pdf
	m.SetBackgroundColor(backgroundColor)
	m.SetBorder(false)
	var column1 uint = 4
	var columnData uint = 8
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text(prop, props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(data, props.Text{
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
