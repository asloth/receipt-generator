package apartment

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Apartment struct {
	number          string
	owner           string
	parking         string
  deposit         string
}

func loadAparmentData (filePath, sheetName string) ([]Apartment, error) {
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
		return nil, err
	}

	cols := []string{}

  ret := []Apartment{}

	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
			fmt.Println("Column information", cols)
		} else {
      ap := Apartment{}
			for j, colCell := range row {
				colCell = strings.TrimSpace(colCell) //el valor de la celda 
        switch j {
        case 0:
          ap.number = colCell
        case 1:
          ap.owner = colCell
        case 2:
          ap.parking = colCell
        case 3:
          ap.deposit = colCell
          break v
        }
			}
		}

	}
	return ret, nil
}
