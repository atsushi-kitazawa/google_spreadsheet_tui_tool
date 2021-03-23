package google_drive

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func DriveSample() {
        b, err := ioutil.ReadFile("drive_credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := google_auth.GetDriveClient(config)

        srv, err := drive.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Drive client: %v", err)
        }

        r, err := srv.Files.List().PageSize(10).
                Fields("nextPageToken, files(id, name)").Do()
        if err != nil {
                log.Fatalf("Unable to retrieve files: %v", err)
        }
        fmt.Println("Files:")
        if len(r.Files) == 0 {
                fmt.Println("No files found.")
        } else {
                for _, i := range r.Files {
                        fmt.Printf("%s (%s)\n", i.Name, i.Id)
                }
        }
}
