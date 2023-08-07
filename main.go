package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-website-downloader",
		Width:  660,
		Height: 600,
        MinWidth: 660,
        MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
        Windows: &windows.Options{
            WebviewIsTransparent: true,
            WindowIsTranslucent: true,
            BackdropType: windows.Acrylic,
        },
        Mac: &mac.Options{
            WebviewIsTransparent: true,
            WindowIsTranslucent:  true,
        },
        Linux: &linux.Options{
            WindowIsTranslucent: true,
        },
        Debug: options.Debug{
            OpenInspectorOnStartup: true,
        },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
