package app

import (
	m "gwd/models"
)

// set static config vars, see startup.go for setting session vars
var Config = m.ConfigType{
	AppDBFileName: "db.json",
	AppName:       "go-website-downloader",
}
