package app

import (
	"encoding/json"
	"fmt"
	"gwd/app/download"
	m "gwd/models"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DownloadButton(r m.ResponseType) {
	fmt.Println("DownloadButton()")

	var db m.DBType
	var adt m.ActivityDataType
	adt.Url = r.Url
	adt.FaviconURL = r.FaviconURL
	adt.SessionID = Config.SessionID

	//----------------------------------------------------------------------------------------
	// Add url to db
	//----------------------------------------------------------------------------------------
	// check if file exists
	if _, err := os.Stat(Config.AppDB); err == nil {

		// read file in
		data, err := os.ReadFile(Config.AppDB)
		if err != nil {
			panic(err)
		} else {

			// parse to struct
			err = json.Unmarshal(data, &db)
			if err != nil {
				panic(err)
			} else {

				// check if url is already being run

				// update data obj
				db.Activity.ActivityCount++
				db.Activity.ActivityData = append(db.Activity.ActivityData, adt)
			}
		}
	} else {
		fmt.Println("File NOT found.")
		db.Activity.ActivityCount = 1
		db.Activity.ActivityData = append(db.Activity.ActivityData, adt)
	}

	// update file
	dJSON, _ := json.Marshal(db)
	err := os.WriteFile(Config.AppDB, dJSON, 0644)
	if err == nil {
		fmt.Println("File written.")
	} else {
		fmt.Println(err)
	}
}

func DownloadSite(r m.ResponseType) error {
	fmt.Println("DownloadSite()")

	db1, err := ReadDB()
	if err != nil {
		fmt.Println(err)
	}

	//----------------------------------------------------------------------------------------
	// download website
	//----------------------------------------------------------------------------------------
	time.Sleep(1 * time.Second)

	u, err := url.Parse(r.Url)
	if err != nil {
		panic(err)
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""

	// create folder names
	s := strings.Replace(u.String()[strings.Index(u.String(), "//")+2:], ".", "_", -1)
	if len(s) > 4 && strings.ToUpper(s[:4]) == "WWW_" {
		s = s[4:]
	}
	currentTime := time.Now()
	t := currentTime.Format("20060102_150405")
	log.Println(s, t, r.FaviconURL)

	download.DownloadManager(r.Url, filepath.Join(db1.Settings.ContentDir, s, t), filepath.Join("/", s, t))

	// // Send an HTTP GET request to the specified URL
	// resp, err := http.Get(r.Url)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	// // Create a new file, or overwrite if it already exists
	// err = os.MkdirAll(filepath.Join(db1.Settings.ContentDir, s, t), 0777)
	// if err != nil {
	// 	return err
	// }
	// file, err := os.Create(filepath.Join(db1.Settings.ContentDir, s, t, "index.html"))
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// // Write the body to the file
	// _, err = file.Write(body)
	// if err != nil {
	// 	return err
	// }

	//----------------------------------------------------------------------------------------
	// get favicon
	//----------------------------------------------------------------------------------------

	if r.FaviconURL != "" {
		// Send an HTTP GET request to the specified URL
		resp, err := http.Get(r.FaviconURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Create a new file, or overwrite if it already exists
		err = os.MkdirAll(filepath.Join(db1.Settings.ContentDir, s, t), 0777)
		if err != nil {
			return err
		}
		file, err := os.Create(filepath.Join(db1.Settings.ContentDir, s, t, "favicon.png"))
		if err != nil {
			return err
		}
		defer file.Close()

		// Write the body to the file
		_, err = file.Write(body)
		if err != nil {
			return err
		}
	}

	//----------------------------------------------------------------------------------------
	// update db
	//----------------------------------------------------------------------------------------
	var db m.DBType
	// check if file exists
	if _, err := os.Stat(Config.AppDB); err == nil {

		// read file in
		data, err := os.ReadFile(Config.AppDB)
		if err != nil {
			return err
		} else {

			// parse to struct
			err = json.Unmarshal(data, &db)
			if err != nil {
				return err
			} else {
				var filteredData m.ActivityType
				o := 0
				for i := 0; i < db.Activity.ActivityCount; i++ {
					if db.Activity.ActivityData[i].Url == r.Url && db.Activity.ActivityData[i].SessionID == Config.SessionID {
						// do not add
					} else {
						filteredData.ActivityData = append(filteredData.ActivityData, db.Activity.ActivityData[i])
						o++
					}
				}
				filteredData.ActivityCount = o

				db.Activity = filteredData

				// update file
				dJSON, _ := json.Marshal(db)
				err := os.WriteFile(Config.AppDB, dJSON, 0644)
				if err == nil {
					fmt.Println("File written.")
				} else {
					fmt.Println(err)
				}
				fmt.Println("Function completed after  1 minute.")
				return nil
			}
		}
	} else {
		return err
	}

}
