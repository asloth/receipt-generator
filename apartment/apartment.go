package apartment

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Apartment struct {
	Number          string
	Owner           string
	Parking         string
  Deposit         string
  ParticipationPercentage string
}

func LoadAparmentData (filePath, sheetName string) ([]Apartment, error) {
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

  out:
	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
			fmt.Println("Column information", cols)
		} else {
      ap := Apartment{}
      inside:
			for j, colCell := range row {
				colCell = strings.TrimSpace(strings.ToUpper(colCell)) //el valor de la celda 
        switch j {
        case 0:
          if len(colCell) == 0 {
						break out
					}
          ap.Number = colCell
        case 1:
          ap.Owner = colCell
        case 2:
          ap.Parking = colCell
        case 3:
          ap.Deposit = colCell
        case 4:
          ap.ParticipationPercentage = colCell
          break inside
        }
			}
      ret = append(ret, ap)
		}
	}
	return ret, nil
}
