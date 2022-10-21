package main

import (
	"embed"
	"mangav3/apps"
	"mangav3/apps/file"
	"mangav3/apps/manga"

	"github.com/wailsapp/wails/v2"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := apps.NewApp()

	//connect sqlite db
	// app.connectDatabaseSqlite("mangadb.db")
	apps.ConnectDatabasePostgres("mangav3")

	//manga instanse
	manga := manga.NewManga()
	freader := file.NewFileReader()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "waislplayground",
		Width:         1024,
		Height:        768,
		Assets:        assets,
		AssetsHandler: file.NewFileLoader(),
		Frameless:     false,
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			WebviewUserDataPath:  app.GetUserDataPath(),
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 24, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			manga,
			freader,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
