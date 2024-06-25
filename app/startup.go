package app

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func StartupActions() {

	//decalre variables
	var appDir string
	var contDir string

	// check OS
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		appDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "GoApps", Config.AppName)
		contDir = filepath.Join(os.Getenv("USERPROFILE"), Config.AppName)
	case "darwin":
		appDir = filepath.Join(os.Getenv("HOME"), "Library", "GoApps", Config.AppName)
		contDir = filepath.Join(os.Getenv("HOME"), Config.AppName)
	case "linux":
		appDir = filepath.Join("/tmp", Config.AppName)
		contDir = filepath.Join(os.Getenv("HOME"), Config.AppName)
	default:
		fmt.Printf("%s.\n", opsys)
		appDir = filepath.Join("/tmp", Config.AppName)
		contDir = appDir
	}

	// set session config vars
	Config.AppDB = filepath.Join(appDir, Config.AppDBFileName)
	Config.AppDir = appDir
	Config.SessionID = getSessionID(10)

	// read database file
	db, err := ReadDB()
	if err != nil {
		fmt.Println("error reading db ReadDB()")
		panic(err)
	}

	// content dir
	if db.Settings.ContentDir == "" {
		db.Settings.ContentDir = contDir
	}
	_, err = os.Stat(db.Settings.ContentDir)
	if err == nil {
		db.Settings.ContentDirExists = true
	} else if os.IsNotExist(err) {
		db.Settings.ContentDirExists = false
	} else {
		db.Settings.ContentDirExists = true
	}

	// create application folder if not exists
	err = os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	// update file
	dJSON, _ := json.Marshal(db)
	err = os.WriteFile(Config.AppDB, dJSON, 0644)
	if err == nil {
		log.Println("File written.")
	} else {
		log.Println(err)
	}

	// remove content dir tmp file
	ShutdownContentDirWebServer()
}

func getSessionID(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	unixTime := time.Now().Unix()
	unixTimeStr := strconv.FormatInt(unixTime, 10)
	return unixTimeStr + "-" + string(s)
}
