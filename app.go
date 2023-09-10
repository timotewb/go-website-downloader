package main

import (
	"context"
	"fmt"

	lib "gwd/lib"
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
}

// Greet returns a greeting for the given name
func (a *App) FindURL(url string) lib.ResponseType {

	// Validate URL
	fmt.Println("--- Validate Start")
	resp, err := lib.ValidateURL(url)
	fmt.Println("--- Validate End")
	if err != nil {
		return resp
	}

	// Verify URL
	fmt.Println("--- Verify")
	resp, page, err := lib.VerifyURL(resp)
	if err != nil {
		return resp
	}

	// Get Favicon
	fmt.Println("--- Favicon")
	resp, err = lib.GetFavicon(resp, page)
	if err != nil {
		return resp
	}

	return resp

}
