package egnyte

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type facility string

// facility names
const (
	VancouverFacility facility = "Vancouver"
	TorontoFacility   facility = "Toronto"
)

// Egnyte RL Inventory All Fields file direct access link prefix and suffix
// Example for RL Inventory All Fields files
// [/Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv]
// [/Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv]
const (
	linkPrefix      = "https://cloudblue.egnyte.com/pubapi/v1/fs-content/Shared/Operations-RL/Daily%20RL%20All%20Fields%20Reports/RL%20Inventory%20All%20Fields_"
	vancouverSuffix = "_Vancouver%2C%20BC%20(RL).csv"
	torontoSuffix   = "_Toronto%2C%20ON.csv"
	// ddmmyyyy format used in the link
	linkDateFormat = "02012006"
)

// Egnyte is used to download files from egnyte.com
type Egnyte struct {
	token             string
	facility          facility
	allFieldsFileLink string
	linkDate          string // format: ddmmyyyy(13072021)
}

// New create a new Egnyte instance
func New(token string, fa facility) *Egnyte {
	if token == "" || fa == "" {
		return nil
	}

	e := &Egnyte{
		token: token,
	}

	err := e.setFacility(fa)
	if err != nil {
		return nil
	}

	e.setInventoryAllFieldsFileLink()

	return e
}

// setFacility set the egnyte facility and calculate the file link
func (e *Egnyte) setFacility(fa facility) error {
	switch fa {
	case VancouverFacility, TorontoFacility:
		e.facility = fa
		return nil
	case "":
		return fmt.Errorf("setFacility: missing facility")
	default:
		return fmt.Errorf("setFacility: facility %v is not supported", fa)
	}
}

// setInventoryAllFieldsFileLink calculate the direct access link to the RL Inventory All Fields file in
// https://cloudblue.egnyte.com/
func (e *Egnyte) setInventoryAllFieldsFileLink() {
	// always download the latest file which was uploaded last night
	// the uploaded date as part of the url link.
	linkDate := time.Now().AddDate(0, 0, -1).Format(linkDateFormat)

	switch e.facility {
	case VancouverFacility:
		// For Vancouver: /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Vancouver, BC (RL).csv
		e.allFieldsFileLink = linkPrefix + linkDate + vancouverSuffix
	case TorontoFacility:
		// /Shared/Operations-RL/Daily RL All Fields Reports/RL Inventory All Fields_13072021_Toronto, ON.csv
		e.allFieldsFileLink = linkPrefix + linkDate + torontoSuffix
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
	req.Header.Add("Authorization", "Bearer "+e.token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return n, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return n, fmt.Errorf("Download: can't download (status code: %v)", response.StatusCode)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return n, fmt.Errorf("Download: can't create file %s: %v", filePath, err)
	}
	defer out.Close()

	return io.Copy(out, response.Body)
}
