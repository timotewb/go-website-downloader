package app

import (
	"fmt"
	m "gwd/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ListGallery() m.ListGalleryType {
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("list gallery")
	fmt.Println("----------------------------------------------------------------------------------------")

	// set variables
	var response m.ListGalleryType
	var faviconLoc string
	dateFmt := "20060102_150405"

	// get settings - content dir
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
	}

	// list folders in content directory and add to response var
	files, err := os.ReadDir(settings.ContentDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			// folder found, get latest favicon location
			siteFiles, err := os.ReadDir(filepath.Join(settings.ContentDir, file.Name()))
			if err != nil {
				fmt.Println(err)
			}
			faviconLoc = ""
			for _, siteFile := range siteFiles {
				if siteFile.IsDir() {
					if faviconLoc == "" {
						faviconLoc = siteFile.Name()
					} else {
						t1, err := time.Parse(dateFmt, faviconLoc)
						if err != nil {
							fmt.Println(err)
						}
						t2, err := time.Parse(dateFmt, siteFile.Name())
						if err != nil {
							fmt.Println(err)
						}
						// if current favicon date is before new date, set faviconLoc to new date
						if t1.Before(t2) {
							faviconLoc = siteFile.Name()
						}
					}
				}
			}
			response = append(response, m.GalleryType{
				SiteName:     strings.Replace(file.Name(), "_", ".", -1),
				SiteLocation: file.Name(),
				Favicon:      "http://localhost:" + strconv.Itoa(int(settings.ContentDirWSPort)) + "/" + file.Name() + "/" + faviconLoc + "/" + "favicon.png",
			})
		}
	}
	// Start webserver
	go RunWebServer()

	// return list of gallery items
	return response
}

func ListGallerySite(siteName string) m.ListGallerySiteType {
	var response m.ListGallerySiteType

	// convert siteName to siteNamePath (from pretty to valid folder path)
	siteNamePath := strings.ReplaceAll(siteName, ".", "_")

	// get settings - content dir
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
	}

	// list folders in content directory and add to response var
	files, err := os.ReadDir(filepath.Join(settings.ContentDir, siteNamePath))
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			response = append(response, m.GallerySiteType{
				DateTime: file.Name(),
				Favicon:  "http://localhost:" + strconv.Itoa(int(settings.ContentDirWSPort)) + "/" + siteNamePath + "/" + file.Name() + "/" + "favicon.png",
			})
		}
	}

	return response
}
