package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// SetColsWidth set the width of one or many columns in the sheetName of the file.
// colsWidth is a map of the column name and width pair
// For Example:
//		error := setColsWidth(f, "Sheet1", map[string]float64{
//			"A": 16.0,
//			"B": 19.0,
//			"C": 40.0,
//			"K": 14.0,
//		})
func SetColsWidth(file *excelize.File, sheetName string, colsWidth map[string]float64) error {
	var err error
	for col, width := range colsWidth {
		err = file.SetColWidth(sheetName, col, col, width)
		if err != nil {
			return err
		}
	}
	return nil
}
