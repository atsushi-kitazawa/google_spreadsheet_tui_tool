package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_drive"
	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_sheet"
)

const credential = "sheet_credentials.json"
const token = "sheet_token.json"

func main() {
	// parse program arguments.
	target := os.Args[1]
	rangeRead := os.Args[2]

	// get target file id from google drive.
	driveFile := google_drive.GetDriveFiles(10)
	var spreadsheetId string
	for _, f := range driveFile {
		if target == f.Name {
			spreadsheetId = f.Id
			break
		}
	}

	if spreadsheetId == "" {
		log.Fatal("target file not found.")
	}

	valRange := google_sheet.ReadSheet(spreadsheetId, rangeRead)

	if len(valRange.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("No, 大項目")
		for _, row := range valRange.Values {
			fmt.Printf("size=%d, %s, %s\n", len(row), row[0], row[1])
		}
	}
}
