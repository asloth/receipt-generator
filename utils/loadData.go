package utils

import (
	"fmt"

	"github.com/asloth/receipt-generator/water"
	"github.com/xuri/excelize/v2"
)

func LoadWaterBuilding(filePath,sheetName string, finalColumn int) (*water.WaterByMonth,error)  {
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
		fmt.Println(err)
		return nil, err
	}

	ret := water.WaterByMonth{}
  ret.Consume = rows[1][1]
  ret.Consumo_rec = rows[2][1]
  ret.Rec_soles = rows[3][1]
  ret.Soles_m3 = rows[4][1]

	return &ret, nil 
}
