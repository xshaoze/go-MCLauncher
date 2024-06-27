package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "myProject",
		Height:    470,
		MinHeight: 470,
		Width:     770,
		MinWidth:  770,

		Frameless: true,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WindowIsTranslucent:               true, // 窗口透明
			DisableFramelessWindowDecorations: true, // 去除窗口装饰
			BackdropType:                      windows.None,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
