# Egnyte

Download Inventory All Fields files from Egnyte.com

## Example
```go
package main

import (
    "fmt"
	"log"
	"os"
	
	"github.com/AngangGuo/lib/egnyte"
)

var (
	vanCSVFileName = "Vancouver.csv"
	torCSVFileName = "Toronto.csv"
)

func main() {
	facilityFilenameList := map[egnyte.Facility]string{
		egnyte.VancouverFacility: vanCSVFileName,
		egnyte.TorontoFacility:   torCSVFileName,
	}
	
	err := downloadFiles(facilityFilenameList)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done.")
}

func downloadFiles(facilityFilenameList map[egnyte.Facility]string) error {
	egnyteTokenName := "EgnyteToken"
	token := os.Getenv(egnyteTokenName)
	if token == "" {
		return fmt.Errorf("can't find the egnyte token from environment variable: %s", egnyteTokenName)
	}

	e := egnyte.New(token)

	for facility, csvName := range facilityFilenameList {
		err := e.SetFacility(facility)
		if err != nil {
			return fmt.Errorf("can't set facility %v", facility)
		}

		log.Printf("downloading Inventory All Fields file for %v from Egnyte.com...\n", facility)

		bytes, err := e.Download(csvName)
		if err != nil {
			return err
		}

		log.Printf("%s: %v bytes downloaded", csvName, bytes)
	}

	return nil
}
```