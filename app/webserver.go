package app

import (
	"context"
	"fmt"
	"gwd/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// flag indicating if handler has been 'handled' or not
var handlerSet bool = false

func RunWebServer() {
	// get settings
	settings, err := GetSettings()
	if err != nil {
		log.Printf("error from RunWebServer(): %v\n", err)
		return
	}

	// define temporary file name
	fileName := filepath.Join(settings.ContentDir, "server_running")

	// check if webserver is already running
	if checkWebServer(settings) {
		return
	} else {

		// create tmp file indicating webserver is running
		file, err := os.Create(fileName)
		if err != nil {
			log.Printf("error from RunWebServer(): os.Create(): - %v\n", err)
			return
		}
		file.Close()

		// Define srv
		srv := &http.Server{
			Addr: ":" + strconv.Itoa(int(settings.ContentDirWSPort)),
		}

		// Channel to receive shutdown signals
		var wg sync.WaitGroup
		defer wg.Done()
		shutdownChan := make(chan struct{})

		// Function to check the shutdown file
		checkForShutdown := func() bool {
			_, err := os.ReadFile(fileName)
			if err == nil || err == io.EOF {
				return false
			}
			return true
		}

		// Start a goroutine to periodically check for the shutdown file
		go func() {
			for {
				time.Sleep(1 * time.Second)
				if checkForShutdown() {
					log.Println("note from RunServer(): Shutting down webserver.")
					shutdownChan <- struct{}{}
					return
				}
			}
		}()

		if !handlerSet {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/" {
					fileName := r.URL.Path[1:]
					filePath := filepath.Join(settings.ContentDir, fileName)
					if _, err := os.Stat(filePath); err != nil {
						http.Error(w, "File not found", http.StatusNotFound)
						return
					}
					http.ServeFile(w, r, filePath)
				} else {
					htmlContent := `<html><body><h1>Hello, World!</h1></body></html>`
					w.Header().Set("Content-Type", "text/html")
					w.Write([]byte(htmlContent))
				}
			})
			handlerSet = true
		}
		go func() {
			// Start the server
			if err := srv.ListenAndServe(); err != nil {
				fmt.Println("Server has been shut down.")
			}
		}()

		// Wait for the server to shut down
		wg.Add(1)
		<-shutdownChan
		fmt.Println("Server shutting down...")
		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("Server forced to shutdown: %v\n", err)
		}
	}
}

func checkWebServer(settings models.SettingsType) bool {
	url := "http://localhost:" + strconv.Itoa(int(settings.ContentDirWSPort))

	// Send a GET request to the server
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("from checkWebServer(): http.Get(): %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("from checkWebServer(): io.ReadAll(): %v\n", err)
		return false
	}

	// Check the status code of the response
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	} else {
		fmt.Printf("Server is not responding or encountered an error. Status Code: %d", resp.StatusCode)
		return false
	}
}

func ShutdownContentDirWebServer() {

	// get settings
	settings, err := GetSettings()
	if err != nil {
		log.Printf("error from RunWebServer(): %v\n", err)
		return
	}

	// define temporary file name
	fileName := filepath.Join(settings.ContentDir, "server_running")

	// Attempt to delete the file
	err = os.Remove(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Failed to delete file: %v", err)
		}
	} else {
		log.Println("File deleted successfully.")
	}
}
