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
func (a *App) FindURL(url string) int {
	lib.ValidateURL(url)
	statusCode, err := lib.VerifyURL(url)
	if err != nil {
		fmt.Println(err)
		return statusCode
	} else{
		fmt.Println(statusCode)
	}
	return statusCode
}
