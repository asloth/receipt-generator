package receipt

import (
	"strings"

	"github.com/asloth/receipt-generator/building"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
)

func ReceiptHeader(pdf *core.Maroto, heightHeader float64, b *building.Building) {
	m := *pdf
	
	var colWidth int = 4

	col1 := col.New(colWidth)
	col1.Add(
		text.New("RECIBO DE MANTENIMIENTO",props.Text{
			Size:  12,
			Style: fontstyle.Bold,
			Align: align.Center,
			Top:   9,
			Color: &props.Color{
				Red: 255,
			},
		}),
		text.New(b.Name, props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Center,
			Top:   15,
		}),
		text.New(b.Address, props.Text{
			Size:  8,
			Style: fontstyle.Bold,
			Align: align.Center,
			Top:   20,
		}),
	)

	m.AddRow(heightHeader, 
		image.NewFromFileCol(
			2,
			"files/molio-logo.jpg",
			props.Rect{	
			Center: true,
			},
		),
		col1,
		image.NewFromFileCol(
			2,
			b.Picture,
			props.Rect{	
			Center: true,
			},
		),
	)
}

func DataOwner(pdf *core.Maroto, backgroundColor props.Color, rowHeight float64, contentSize float64, prop1, data1, prop2, data2 string) {
	m := *pdf
	var column1 int = 4
	var columnData int = 2
	m.AddRow(9, 
		text.NewCol(column1,strings.ToUpper(prop1),
			props.Text{
				Size:            8,
				Align:           align.Center,
				Style:           fontstyle.Bold,
				VerticalPadding: 0,
				Top:             4,
			},
		),
		text.NewCol(columnData,data1, props.Text{
			Top: 4,
		}),
		text.NewCol(column1,strings.ToUpper(prop2), props.Text{
			Size:            8,
			Align:           align.Center,
			Style:           fontstyle.Bold,
			VerticalPadding: 0,
			Top:             4,
		}),
		text.NewCol(2,data2, props.Text{
			Top: 4,
		}),
	)
}

func PayInfo(pdf *pdf.Maroto, colorMolio props.Color, b *building.Building) {
	m := *pdf

	headers := []string{"BANCO", "CUENTA BANCARIA", "TITULAR DE CUENTA"}
	contents := [][]string{
		{b.Bank, b.BankAccount, b.BankAccountOwner},
	}

	m.TableList(headers, contents,  TableList{
		HeaderProp: props.TableListContent{
			Family:    fontfamily.Arial,
			Style:     fontstyle.Bold,
			Size:      11.0,
			GridSizes: []uint{4, 4, 4},
		},
		ContentProp: props.TableListContent{
			Family:    fontfamily.Courier,
			Style:     fontstyle.Normal,
			Size:      10.0,
			GridSizes: []uint{4, 4, 4},
		},
		Align: align.Center,
		HeaderContentSpace:     0.01,
		VerticalContentPadding: 4.0,
		AlternatedBackground: &props.Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		},
	})
}

func SubHeader(pdf *pdf.Maroto, colorMolio props.Color, subtitulo string) {
	m := *pdf
	m.SetBorder(true)
	m.SetBackgroundColor(colorMolio)
	m.AddRow(7, func() {
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

func Footer(pdf *pdf.Maroto, backgroundColor props.Color, contentSize float64) {
	m := *pdf
	m.SetBackgroundColor(backgroundColor)
	m.AddRow(10, func() {
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

func Resumen(pdf *pdf.Maroto, backgroundColor props.Color, contentSize float64, field, amount string) {
	m := *pdf
	m.AddRow(7, func() {
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

func ApartmentData(pdf *pdf.Maroto, backgroundColor props.Color, contentSize float64, field, value string) {

	m := *pdf
	m.SetBackgroundColor(backgroundColor)

	m.AddRow(7, func() {
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
