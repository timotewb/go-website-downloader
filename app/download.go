package app

import (
	"encoding/json"
	"fmt"
	m "gwd/models"
	"net/url"
	"os"
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

	//----------------------------------------------------------------------------------------
	// download website
	//----------------------------------------------------------------------------------------
	time.Sleep(5 * time.Second)

	u, err := url.Parse(r.Url)
	if err != nil {
		panic(err)
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""

	s := strings.Replace(u.String()[strings.Index(u.String(), "//"):], ".", "_", -1)
	fmt.Println(s)

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
	// file, err := os.Create("test_site.html")
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
