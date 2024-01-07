package utils

import (
	"fmt"
	"strconv"

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
  
  real_consumption, err := strconv.ParseFloat(rows[1][1], 64)
  recep_consumption, err := strconv.ParseFloat(rows[2][1], 64)
  total_charge, err := strconv.ParseFloat(rows[3][1], 64)
  charge_per_m3, err := strconv.ParseFloat(rows[4][1], 64)

	ret := water.WaterByMonth{}
  ret.Consume = fmt.Sprintf("%.2f", real_consumption)
  ret.Consumo_rec = fmt.Sprintf("%.2f", recep_consumption)
  ret.Rec_soles = fmt.Sprintf("S/. %.2f", total_charge)
  ret.Soles_m3 = fmt.Sprintf("S/. %.2f", charge_per_m3)

	return &ret, nil 
}
