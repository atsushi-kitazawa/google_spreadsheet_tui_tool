package main

import (
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

	google_sheet.ReadSheet(spreadsheetId, rangeRead)
}
