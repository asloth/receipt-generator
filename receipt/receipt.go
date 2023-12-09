package receipt

import (
	"strings"

	"github.com/asloth/receipt-generator/building"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func ReceiptHeader(pdf *pdf.Maroto, heightHeader float64, b *building.Building) {
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
			m.Text(b.Name, props.Text{
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
				Top:   15,
			})
			m.Text(b.Address, props.Text{
				Size:  8,
				Style: consts.Bold,
				Align: consts.Center,
				Top:   20,
			})
		})
		m.Col(colWidth, func() {
			_ = m.FileImage(b.Picture, props.Rect{
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
			m.Text(strings.ToUpper(prop1), props.Text{
				Size:            8,
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
			m.Text(strings.ToUpper(prop2), props.Text{
				Size:            8,
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

func PayInfo(pdf *pdf.Maroto, colorMolio color.Color, b *building.Building) {
	m := *pdf

	headers := []string{"BANCO", "CUENTA BANCARIA", "TITULAR DE CUENTA"}
	contents := [][]string{
		{b.Bank, b.BankAccount, b.BankAccountOwner},
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

func Resumen(pdf *pdf.Maroto, backgroundColor color.Color, contentSize float64, field, amount string) {
	m := *pdf
	m.Row(7, func() {
		m.Col(10, func() {
			m.Text(field,
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Left,
				})
		})
		m.Col(2, func() {
			m.Text(amount,
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Center,
				})
		})
	})
}

func ApartmentData(pdf *pdf.Maroto, backgroundColor color.Color, contentSize float64, field, value string) {

	m := *pdf
  m.SetBackgroundColor(backgroundColor)

	m.Row(7, func() {
		m.Col(4, func() {
			m.Text(field,
				props.Text{
					Size:  contentSize,
					Style: consts.Bold,
					Align: consts.Right,
				})
		})
		m.Col(8, func() {
			m.Text(value,
				props.Text{
					Size:  contentSize,
					Style: consts.Normal,
					Align: consts.Left,          
				})
		})
	})
}
