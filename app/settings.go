package app

import (
	"context"
	"encoding/json"
	"log"
	"os"

	r "github.com/wailsapp/wails/v2/pkg/runtime"
)

func UpdateContentDir(ctx context.Context) {
	// read database file
	db, err := ReadDB()
	if err != nil {
		log.Printf("error from UpdateContentDir():")
		log.Printf("    ReadDB(): %v", err)
		return
	}

	dir, err := r.OpenDirectoryDialog(ctx, r.OpenDialogOptions{
		DefaultDirectory:     db.Settings.ContentDir,
		Title:                "current-content-dir",
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		log.Printf("error from UpdateContentDir():")
		log.Printf("    OpenDirectoryDialog(): %v", err)
		return
	}

	if dir != "" {
		// update database
		db.Settings.ContentDir = dir

		// update file
		dJSON, _ := json.Marshal(db)
		err = os.WriteFile(Config.AppDB, dJSON, 0644)
		if err != nil {
			log.Printf("error from UpdateContentDir():")
			log.Printf("    os.WriteFile(): %v", err)
			return
		}
	}
}
