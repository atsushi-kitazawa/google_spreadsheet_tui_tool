package google_drive

import (
	_ "fmt"
	"io/ioutil"
	"log"

	"github.com/atsushi-kitazawa/google_spreadsheet_tui_tool/google_auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const credential = "drive_credentials.json"
const token = "drive_token.json"

type DriveFile struct {
    Name string
    Id string
}

func getDriveService() *drive.Service {
        b, err := ioutil.ReadFile(credential)
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := google_auth.GetClient(config,token)

        srv, err := drive.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Drive client: %v", err)
        }
	return srv
}

func GetDriveFiles(pageSize int64) []DriveFile {
	srv := getDriveService()

	r, err := srv.Files.List().PageSize(pageSize).
                Fields("nextPageToken, files(id, name)").Do()
        if err != nil {
                log.Fatalf("Unable to retrieve files: %v", err)
        }

	fileList := make([]DriveFile, 0)
        if len(r.Files) == 0 {
                log.Fatal("No files found.")
        } else {
                for _, i := range r.Files {
                        //fmt.Printf("%s (%s)\n", i.Name, i.Id)
			fileList = append(fileList, DriveFile{Name:i.Name, Id:i.Id, })
                }
        }
	return fileList
}
