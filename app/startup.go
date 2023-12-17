package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func StartupActions(){

	//decalre variables
	var appDir string

	// check OS
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		appDir = filepath.Join("C:\\Temp", Config.AppName)
	case "darwin":
		appDir = filepath.Join(os.Getenv("HOME"),"Library/GoApps/",Config.AppName)
	case "linux":
		appDir = filepath.Join("/tmp", Config.AppName)
	default:
		fmt.Printf("%s.\n", opsys)
		appDir = filepath.Join("/tmp", Config.AppName)
	}

	// run
	actions(appDir)
}

func actions(appDir string){
	// create Library folder if not exists
	err := os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	// set session vars
	Config.AppDB = filepath.Join(appDir, Config.AppDBFileName)
	Config.AppDir = appDir

}