package app

import (
	"encoding/json"
	m "gwd/models"
	"os"
)

// set static config vars, see startup.go for setting session vars
var Config = m.ConfigType{
	AppDBFileName: "db.json",
	AppName:       "go-website-downloader",
}

func ReadDB() (m.DBType, error) {
	var db m.DBType
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
				return db, err
			}
		}
	} else {
		return db, err
	}
}

func GetSettings() (m.SettingsType, error) {
	db, err := ReadDB()
	return db.Settings, err
}
