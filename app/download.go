package app

import (
	"encoding/json"
	"fmt"
	m "gwd/models"
	"os"
)

func DownloadButton(r m.ResponseType) {
	fmt.Println("DownloadButton()")

	var db m.DBType
	var adt m.ActivityDataType
	adt.Url = r.Url
	adt.FaviconURL = r.FaviconURL

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
				fmt.Println(db)
			}
		}
	} else{
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