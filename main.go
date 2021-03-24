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
const targetArg = 1
const targetRangeArg = 2
const pageSize = 10

func main() {
	// parse program arguments.
	target := os.Args[targetArg]
	rangeRead := os.Args[targetRangeArg]

	// get target file id from google drive.
	driveFile := google_drive.GetDriveFiles(pageSize)
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

	// get target sheet value range.
	valRange := google_sheet.ReadSheet(spreadsheetId, rangeRead)
	if len(valRange.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range valRange.Values {
			fmt.Printf("size=%d, %s, %s\n", len(row), row[0], row[1])
		}
	}
}
