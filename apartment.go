package main

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	owner := "BENITO JAVIER, SANTOS MEDINA"

	var heightHeader float64 = 30
	var contentSize float64 = 10
	var rowHeight float64 = 7
	colorMolio := color.Color{
		Red:   96,
		Green: 237,
		Blue:  26,
	}
	backgroundColor := color.NewWhite()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	// Header
	m.Row(heightHeader, func() {
		m.Col(4, func() {
			_ = m.FileImage("files/molio-logo.jpg", props.Rect{

				Center: true,
			})
		})
		m.Col(4, func() {
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
		m.Col(4, func() {
			_ = m.FileImage("files/parque-roma-logo.jpg", props.Rect{
				Center: true,
			})
		})
	})

	// tabla inicial
	headers := []string{"TIPO CUOTA", "F. EMISION", "F. VCTO.", "PERIODO", "N. RECIBO"}
	contents := [][]string{
		{"Content1", "Content2", "Content2", "Content2", "Content2"},
	}
	m.Line(10)
	m.SetBorder(true)
	m.SetBackgroundColor(colorMolio)

	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Bold,
			Size:      11.0,
			GridSizes: []uint{3, 2, 2, 3, 2},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			Size:      10.0,
			GridSizes: []uint{3, 2, 2, 3, 2},
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
	m.SetBackgroundColor(colorMolio)
	m.Row(7, func() {
		m.Col(12, func() {
			m.Text("DATOS DEL PROPIETARIO/INQUILINO",
				props.Text{
					Size:  12,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})
	m.SetBackgroundColor(backgroundColor)
	m.SetBorder(false)
	var column1 uint = 4
	var columnData uint = 8
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("NOMBRE: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("DEPARTAMENTO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("CODIGO BANCO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("ESTACIONAMIENTO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("AREA DEPARTAMENTO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("AREA TOTAL DEL EDIFICIO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})
	m.Row(rowHeight, func() {
		m.Col(column1, func() {
			m.Text("TOTAL PRESUPUESTO: ", props.Text{
				Size:            contentSize,
				Align:           consts.Right,
				Style:           consts.Bold,
				VerticalPadding: 3,
				Top:             1,
			})
		})
		m.Col(columnData, func() {
			m.Text(owner, props.Text{
				Top: 1,
			})
		})
	})

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
			m.Text("MANTENIMIENTO ( 0.5841 % ) x ( S/ 28,974.00 )",
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Left,
				})
		})
		m.Col(2, func() {
			m.Text("S/. 169.24",
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})

	m.OutputFileAndClose("maroto.pdf")
}
