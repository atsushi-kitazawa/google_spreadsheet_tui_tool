package google_sheet

import (
	"io/ioutil"
	"log"

	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const credential = "sheet_credentials.json"
const token = "sheet_token.json"

func ReadSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	b, err := ioutil.ReadFile(credential)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := google_auth.GetClient(config, token)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp
}
