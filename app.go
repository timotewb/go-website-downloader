package main

import (
	"context"
	"fmt"
	"gwd/lib"
	"net/url"
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
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Validate URL
func (a *App) ValidateURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Check URL
func (a *App) GetFavicon(url string) string {
	r, err := lib.GetFavicon(url)
    if err != nil {
        fmt.Println("Error:", err)
        panic(err)
    }
	return r
}