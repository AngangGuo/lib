package egnyte

import (
	"fmt"
	"strings"
	"time"
)

// facility names
const (
	VancouverFacility = "Vancouver"
	TorontoFacility   = "Toronto"
)

// RL Inventory All Fields file download link
// Example link for Vancouver: /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv
// Example link for Toronto: /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv
const (
	linkPrefix = "https://cloudblue.egnyte.com/pubapi/v1/fs-content/Shared/Operations-RL/Daily%20RL%20All%20Fields%20Reports/RL%20Inventory%20All%20Fields_"
	// ddmmyyyy format used in the link
	linkDateFormat  = "02012006"
	vancouverSuffix = "_Vancouver%2C%20BC%20(RL).csv"
	torontoSuffix   = "_Toronto%2C%20ON.csv"
)

// getFileLink return the direct api link of the latest RL Inventory All Fields file of the facility
func getFileLink(facility string) (string, error) {
	ddmmyyyyDate := time.Now().AddDate(0, 0, -1).Format(linkDateFormat)

	var link string
	switch facility {
	case VancouverFacility:
		// For Vancouver: /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv
		link = linkPrefix + ddmmyyyyDate + vancouverSuffix
		return link, nil
	case TorontoFacility:
		// /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv
		link = linkPrefix + ddmmyyyyDate + torontoSuffix
		return link, nil
	default:
		link = ""
		return link, fmt.Errorf("getFileLink: facility %s is not supported", facility)
	}
}

// GetSavedFileName returns the .csv file name according to facility names
// facility can only be Vancouver or Toronto.
// the file name is used to save the Daily RL All Fields Reports
// Extract this function to coordinate the download program and the report program
func GetSavedFileName(facility string) (string, error) {
	facility = strings.Title(strings.TrimSpace(facility))

	switch facility {
	case VancouverFacility, TorontoFacility:
		return fmt.Sprintf("%s.csv", facility), nil
	default:
		return "", fmt.Errorf("GetSavedFileName: wrong facility %s", facility)
	}
}
