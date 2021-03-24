package google_sheet

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const credential = "sheet_credentials.json"
const token = "sheet_token.json"
const target = "google spreadsheet tui tool"

func ReadSheet(spreadsheetId string) {
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

	readRange := "管理!B5:I30"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("No, 大項目")
		for _, row := range resp.Values {
			fmt.Printf("size=%d, %s, %s\n", len(row), row[0], row[1])
		}
	}
}
