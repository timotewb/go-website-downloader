package app

import (
	"encoding/json"
	"fmt"
	m "gwd/models"
	"os"
)

func RemoveStaleActivity(url string, sessionID string) error {
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
					if db.Activity.ActivityData[i].Url == url && db.Activity.ActivityData[i].SessionID == sessionID {
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
				return nil
			}
		}
	} else {
		return err
	}
}

func CheckActivity() (m.CheckActivityType, error) {

	var db m.DBType
	var cat m.CheckActivityType

	// check if file exists
	if _, err := os.Stat(Config.AppDB); err == nil {

		// read file in
		data, err := os.ReadFile(Config.AppDB)
		if err != nil {
			return cat, err
		} else {

			// parse to struct
			err = json.Unmarshal(data, &db)
			if err != nil {
				return cat, err
			} else {
				cat.JobCount = db.Activity.ActivityCount
				cat.ActivityData = db.Activity.ActivityData
			}
		}
	} else {
		return cat, nil
	}
	// flag if activity item was NOT created this session - stale
	if cat.JobCount >= 1 {
		for i := 0; i < cat.JobCount; i++ {
			if cat.ActivityData[i].SessionID == Config.SessionID {
				cat.ActivityData[i].StaleFlag = false
			} else {
				cat.ActivityData[i].StaleFlag = true
			}
		}
	}
	return cat, nil

}
