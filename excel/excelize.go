package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
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

// GetTitleColList returns the list of title and column number pairs
// tableTitleList: list of table titles from a sheet
// namesToSearch: the list of names to search in the tableTitles
// returns error if any one of the names can't be found from the title
//
// Note: the returned column number is 0 based
func GetTitleColList(tableTitleList, namesToSearch []string) (map[string]int, error) {
	// initialize title name:column number pair
	var nameCol = map[string]int{}
	for _, name := range namesToSearch {
		// init to any value, just to make a map with the title name
		nameCol[name] = -1
	}

	// set the actual column number for the title
	for i, j := range tableTitleList {
		if _, ok := nameCol[j]; !ok {
			continue
		}
		nameCol[j] = i
	}

	for k, v := range nameCol {
		if v == -1 {
			return nil, fmt.Errorf("can't find the column with title %s", k)
		}
	}
	return nameCol, nil
}
