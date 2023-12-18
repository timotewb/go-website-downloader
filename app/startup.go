package app

import (
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

	// check OS
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		appDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "GoApps", Config.AppName)
	case "darwin":
		appDir = filepath.Join(os.Getenv("HOME"), "Library", "GoApps", Config.AppName)
	case "linux":
		appDir = filepath.Join("/tmp", Config.AppName)
	default:
		fmt.Printf("%s.\n", opsys)
		appDir = filepath.Join("/tmp", Config.AppName)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("appDir: ", appDir)
	fmt.Println("")
	fmt.Println("")

	// run
	actions(appDir)
}

func actions(appDir string) {
	// create application folder if not exists
	err := os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	// set session config vars
	Config.AppDB = filepath.Join(appDir, Config.AppDBFileName)
	Config.AppDir = appDir
	Config.SessionID = getSessionID(10)

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
