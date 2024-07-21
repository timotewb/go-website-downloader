package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func DownloadManager(url string, rootDir string, rootWSDir string) {

	var dm downloadManagerType
	dm.UserURL = url
	domain, err := getRootDomain(url)
	if err != nil {
		log.Printf("error from getRootDomain(): %v", err)
	}
	dm.Domain = domain
	rootURL, err := getRootURL(url)
	if err != nil {
		log.Printf("error from getRootURL(): %v", err)
	}
	dm.RootURL = rootURL
	dm.RootDir = rootDir
	dm.RootWSDir = rootWSDir
	dm.LogFile = filepath.Join(rootDir, "job_log.log")

	// define linkType for root url
	var lt linkType
	lt.GetURL = rootURL
	lt.Kind = "page"
	lt.SaveDir = filepath.Join(dm.RootDir, lt.Kind, "index.html")
	lt.ValNew = filepath.Join(dm.RootWSDir, lt.Kind, "index.html")
	lt.ValOriginal = "/"
	lt.WrittenOut = false
	lt.Attr = "href"
	lt.Data = "a"

	dm.Links = append(dm.Links, lt)

	downloadURL(lt, &dm)
}

func downloadURL(lt linkType, dm *downloadManagerType) {
	log.Printf("downloadURL('%v') with kind='%v'\n", lt.GetURL, lt.Kind)
	err := updateLogFile(dm.LogFile, fmt.Sprintf("Kind: %v, GetURL: %v", lt.Kind, lt.GetURL))
	if err != nil {
		log.Printf("err from updateLogFile(): %v", err)
	}

	// download url and parse
	resp, err := http.Get(lt.GetURL)
	if err != nil {
		log.Printf("err from http.Get(): %v", err)
	}
	defer resp.Body.Close()

	// if kind is page, search for links
	if lt.Kind == "page" {
		doc, _ := html.Parse(resp.Body)
		// define crawler and crawl
		var crawler func(*html.Node)
		crawler = func(node *html.Node) {
			if node.Type == html.ElementNode {
				applyActions(node, dm)
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				crawler(child)
			}
		}
		crawler(doc)
		writeOutPage(doc, lt, dm)
	} else if lt.Kind == "resource" {
		// create dirs
		dirName := filepath.Dir(lt.SaveDir)
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Printf("Error creating directory: %v\n", err)
		}
		// Create a file where the downloaded content will be saved
		out, err := os.Create(lt.SaveDir)
		if err != nil {
			log.Printf("Error creating file: %v\n", err)
		}
		defer out.Close()

		// Copy the response body to the file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Printf("Error writing file: %v\n", err)
		}
		updateDMLT(lt, dm)
	}

	// depth +1 - this could cause some funky things with dm being pased around and updated in nested calls
	for _, l := range dm.Links {
		if !l.WrittenOut {
			downloadURL(l, dm)
		}
	}
}
