package main

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	var rowHeight float64 = 30
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.Row(rowHeight, func() {
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

	m.OutputFileAndClose("maroto.pdf")
}
