package egnyte

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Egnyte
type Egnyte struct {
	Token             string
	facility          string
	allFieldsFileLink string
	linkDate          string
}

// facility names
const (
	VancouverFacility = "Vancouver"
	TorontoFacility   = "Toronto"
)

// Example for RL Inventory All Fields files
// [/Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv]
// [/Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv]
const (
	linkPrefix = "https://cloudblue.egnyte.com/pubapi/v1/fs-content/Shared/Operations-RL/Daily%20RL%20All%20Fields%20Reports/RL%20Inventory%20All%20Fields_"
	// ddmmyyyy format used in the link
	linkDateFormat  = "02012006"
	vancouverSuffix = "_Vancouver%2C%20BC%20(RL).csv"
	torontoSuffix   = "_Toronto%2C%20ON.csv"
)

// New create a new Egnyte instance
// set EgnyteToken="69y...ctp
// egnyteTokenName := "EgnyteToken"
// token := os.Getenv(egnyteTokenName)
// if token == "" {
// 	return fmt.Errorf("can't find the egnyte token from environment variable: %s", egnyteTokenName)
// }
// egnyte := Egnyte.New(token)
func New(token string) *Egnyte {
	return &Egnyte{
		Token: token,
		// always download the latest file
		// The latest uploaded file on website is from yesterday
		linkDate: time.Now().AddDate(0, 0, -1).Format(linkDateFormat),
	}
}

func (e *Egnyte) SetFacility(facility string) error {
	switch facility {
	case VancouverFacility, TorontoFacility:
		e.facility = facility
		e.setInventoryAllFieldsFileLink()
		return nil
	default:
		return fmt.Errorf("setFacility: facility %s is not supported", facility)
	}
}

// getInventoryAllFieldsFileLink return the direct api link of the latest RL Inventory All Fields file of the facility
func (e *Egnyte) setInventoryAllFieldsFileLink() {
	switch e.facility {
	case VancouverFacility:
		// For Vancouver: /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv
		e.allFieldsFileLink = linkPrefix + e.linkDate + vancouverSuffix
	case TorontoFacility:
		// /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv
		e.allFieldsFileLink = linkPrefix + e.linkDate + torontoSuffix
		//default:
		//	e.allFieldsFileLink = ""
		//	return fmt.Errorf("setInventoryAllFieldsFileLink: facility %s is not supported", facility)
	}
}

// Download will download the Inventory All Fields file(.csv) from Egnyte and write the file to filePath
func (e *Egnyte) Download(filePath string) (int64, error) {
	// number of bytes downloaded
	var n int64

	req, err := http.NewRequest("GET", e.allFieldsFileLink, nil)
	if err != nil {
		return n, err
	}
	req.Header.Add("Authorization", "Bearer "+e.Token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return n, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return n, fmt.Errorf("wrong status code: %v", response.StatusCode)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return n, err
	}
	defer out.Close()

	return io.Copy(out, response.Body)
}
