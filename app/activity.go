package app

import (
	"encoding/json"
	m "gwd/models"
	"os"
)

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
