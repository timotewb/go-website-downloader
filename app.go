package main

import (
	"context"
	"fmt"

	app "gwd/app"
	lib "gwd/lib"
	m "gwd/models"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	app.StartupActions()
}

// Greet returns a greeting for the given name
func (a *App) FindURL(url string) lib.ResponseType {

	// Validate URL
	fmt.Println("--- Validate")
	resp, err := lib.ValidateURL(url)
	if err != nil {
		return resp
	}

	// Verify URL
	fmt.Println("--- Verify")
	resp, page, err := lib.VerifyURL(resp)
	if err != nil {
		return resp
	}
	defer page.Body.Close()

	// Get Favicon
	fmt.Println("--- Favicon")
	resp, err = lib.GetFavicon(resp, page)
	if err != nil {
		return resp
	}
	return resp

}

func (a *App) GetSite(r m.ResponseType) {
	app.DownloadButton(r)
	go app.DownloadSite(r)
}

func (a *App) CheckActivity() m.CheckActivityType {
	resp, err := app.CheckActivity()
	if err != nil {
		return resp
	}
	return resp
}

func (a *App) RemoveStaleActivity(url string, sessionID string) {
	app.RemoveStaleActivity(url, sessionID)
}

func (a *App) GetSettings() m.SettingsType {
	resp, err := app.ReadDB()
	if err != nil {
		return resp.Settings
	}
	return resp.Settings
}

func (a *App) ListGallery() m.ListGalleryType {
	return app.ListGallery()
}

func (a *App) ShutdownContentDirWebServer() {
	app.ShutdownContentDirWebServer()
}

func (a *App) ListGallerySite(siteName string) m.ListGallerySiteType {
	return app.ListGallerySite(siteName)
}
