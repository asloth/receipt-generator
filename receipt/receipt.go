package receipt

import (
	"strings"

	"github.com/asloth/receipt-generator/building"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func ReceiptHeader(pdf *core.Maroto, heightHeader float64, b *building.Building) {
	m := *pdf
	
	var colWidth int = 6

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
			3,
			"files/molio-logo.jpg",
			props.Rect{	
				Center: true,
				Percent: 95,
				
			},
		),
		col1,
		image.NewFromFileCol(
			3,
			b.Picture,
			props.Rect{	
				Center: true,
				Percent: 90,
			},
		),
	)
}

func DataOwner(pdf *core.Maroto, rowHeight float64, contentSize float64, prop1, data1, prop2, data2 string) {
	m := *pdf
	var column1 int = 4
	var columnData int = 2
	m.AddRow(9, 
		text.NewCol(column1,strings.ToUpper(prop1),
			props.Text{
				Size:            8,
				Align:           align.Left,
				Style:           fontstyle.Bold,
				Top:             2,
			},
		),
		text.NewCol(columnData,data1, props.Text{
			Top: 2,
		}),
		text.NewCol(column1,strings.ToUpper(prop2), props.Text{
			Size:            8,
			Align:           align.Left,
			Style:           fontstyle.Bold,
			Top:             2,
		}),
		text.NewCol(2,data2, props.Text{
			Top: 2,
		}),
	)
}

func PayInfo(pdf *core.Maroto, headerColStyle *props.Cell, contentColStyle *props.Cell, b *building.Building) {
	m := *pdf
	headerTextStyle := props.Text{
		Size: 11,
		Style: fontstyle.Bold,
		Align: align.Center,
		Family: fontfamily.Arial,
		Top: 2,
	}
	headers := []core.Col{
		text.NewCol(4,"BANCO",headerTextStyle).WithStyle(headerColStyle),
		text.NewCol(4,"CUENTA BANCARIA",headerTextStyle).WithStyle(headerColStyle),
		text.NewCol(4,"TITULAR DE CUENTA",headerTextStyle).WithStyle(headerColStyle),
	}
	
	contentTextStyle := props.Text{
		Family:    fontfamily.Courier,
		Style:     fontstyle.Normal,
		Top: 1,
		Size:      10.0,
		Align: align.Center,
	}
	contents := []core.Col{
		text.NewCol(4,b.Bank,contentTextStyle).WithStyle(contentColStyle),
		text.NewCol(4,b.BankAccount,contentTextStyle).WithStyle(contentColStyle),
		text.NewCol(4,b.BankAccountOwner,contentTextStyle).WithStyle(contentColStyle),
	}
	m.AddRow(8, headers...)
	m.AddRow(12, contents...)
}

func SubHeader(pdf *core.Maroto, subtitulo string, colStyleHeader *props.Cell) {
	m := *pdf
	m.AddRow(7,
		text.NewCol(12,subtitulo,
			props.Text{
				Size:  12,
				Style: fontstyle.Bold,
				Top: 0.5,
				Align: align.Center,
			}).WithStyle(colStyleHeader),
	)
}

func Footer(pdf *core.Maroto, backgroundColor props.Color, contentSize float64) {
	m := *pdf
	m.AddRow(10, 
			text.NewCol(12,"1. Se deja por escrito que el incumplimiento del pago esta sujeto a mora.",
				props.Text{
					Size:  contentSize,
					Align: align.Left,
				}),
			text.NewCol(12,"2. El Propietario autoriza el corte de suministro de agua por incumplimiento de pago.",
				props.Text{
					Size:  contentSize,
					Align: align.Left,
					Top:   5,
				}),
	)
}

func Resumen(pdf *core.Maroto, colStyleCenterContent *props.Cell,contentSize float64, field, amount string) {
	m := *pdf
	m.AddRow(7, 
		text.NewCol(10,field,
			props.Text{
				Size:  contentSize,
				Style: fontstyle.Bold,
				Top: 1,
				Align: align.Left,
			}).WithStyle(colStyleCenterContent),
		text.NewCol(2,amount,
			props.Text{
				Size:  contentSize,
				Style: fontstyle.Bold,
				Top: 1,
				Align: align.Center,
			}).WithStyle(colStyleCenterContent),
	)
}

func ApartmentData(pdf *core.Maroto, colStyleCenterContent *props.Cell, contentSize float64, field, value string){
	m := *pdf

	m.AddRow(7, 
		text.NewCol(4,field,
			props.Text{
				Size:  contentSize,
				Style: fontstyle.Bold,
				Align: align.Right,
			}).WithStyle(colStyleCenterContent),
		text.NewCol(8,value,
			props.Text{
				Size:  contentSize,
				Left: 2,
				Style: fontstyle.Normal,
				Align: align.Left,
			}).WithStyle(colStyleCenterContent),
	)
}
