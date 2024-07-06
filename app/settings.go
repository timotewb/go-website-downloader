package app

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
		Title:                "Select content folder",
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		log.Printf("error from UpdateContentDir():")
		log.Printf("    OpenDirectoryDialog(): %v", err)
		if strings.Contains(err.Error(), "default directory") {
			dir, err = r.OpenDirectoryDialog(ctx, r.OpenDialogOptions{
				Title:                "Select content folder",
				CanCreateDirectories: true,
				ShowHiddenFiles:      true,
			})
			if err != nil {
				log.Printf("error from UpdateContentDir():")
				log.Printf("    OpenDirectoryDialog() no default: %v", err)
				return
			}
		}
	}

	log.Printf("message from UpdateContentDir(): dir provided %v", dir)
	if dir != "" {
		fileInfo, err := os.Lstat(db.Settings.ContentDir)
		if err != nil {
			log.Printf("error from UpdateContentDir():")
			log.Printf("    os.Lstat(): %v", err)
			log.Println("    fileInfo.Mode() set to 0777")

			// create dir
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				log.Printf("error from UpdateContentDir():")
				log.Printf("    os.MkdirAll() with 0777: %v", err)
				return
			}
			log.Println("message from UpdateContentDir(): dir created")
		} else {
			err = os.MkdirAll(dir, fileInfo.Mode())
			if err != nil {
				log.Printf("error from UpdateContentDir():")
				log.Printf("    os.MkdirAll() with fileInfo.Mode(): %v", err)
				return
			}
			log.Println("message from UpdateContentDir(): dir created")
		}

		// check if current dir contains content
		var contentFound bool = false
		folders, err := filepath.Glob(filepath.Join(db.Settings.ContentDir, "*"))
		if err != nil {
			log.Printf("error from UpdateContentDir():")
			log.Printf("    filepath.Glob(): %v", err)
			return
		}
		for _, folder := range folders {
			_, err := os.Stat(folder)
			if err == nil {
				contentFound = true
				break
			}
		}
		if contentFound {
			answer, err := r.MessageDialog(ctx, r.MessageDialogOptions{
				Type:          runtime.QuestionDialog,
				Title:         "Move content",
				Message:       "Content found in existing directoy.\nDo you want to move it to the new directory?",
				DefaultButton: "Yes",
				Buttons:       []string{"No", "Yes"},
			})
			if err != nil {
				log.Printf("error from UpdateContentDir():")
				log.Printf("    r.MessageDialog(): %v", err)
			}
			if answer == "Yes" {
				err = moveFiles(db.Settings.ContentDir, dir)
				if err != nil {
					log.Printf("error from UpdateContentDir():")
					log.Printf("    moveFiles(): %v", err)
				}

			}
		}

		// update database with new directory
		if strings.HasSuffix(dir, Config.AppName) {
			db.Settings.ContentDir = dir
		} else {
			db.Settings.ContentDir = filepath.Join(dir, Config.AppName)
		}

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

func moveFiles(src string, dst string) error {
	// Use os.Lstat to get FileInfo
	fileInfo, err := os.Lstat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, fileInfo.Mode())
	if err != nil {
		return err
	}

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Create destination file
			dstPath := filepath.Join(dst, filepath.Base(path))
			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			// Copy contents of source file to destination file
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func UpdatePortNumber(port int16) {
	log.Printf("Update port number: %v", port)
	// read database file
	db, err := ReadDB()
	if err != nil {
		log.Printf("error from UpdateContentDir():")
		log.Printf("    ReadDB(): %v", err)
		return
	}

	// update port
	db.Settings.ContentDirWSPort = port

	// update file
	dJSON, _ := json.Marshal(db)
	err = os.WriteFile(Config.AppDB, dJSON, 0644)
	if err != nil {
		log.Printf("error from UpdateContentDir():")
		log.Printf("    os.WriteFile(): %v", err)
		return
	}
}
