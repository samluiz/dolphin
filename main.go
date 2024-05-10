package main

import (
	"dolphin/backend/db"
	"dolphin/backend/earning"
	"dolphin/backend/expense"
	"dolphin/backend/expense/category"
	"dolphin/backend/profile"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var Version string

func main() {

	db.Version = Version

	db, err := db.Connect()

	fmt.Println("version: ", Version)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	earningRepository := earning.NewRepository(db)
	earningService := earning.NewService(earningRepository)

	expenseRepository := expense.NewRepository(db)
	expenseService := expense.NewService(expenseRepository)

	profileRepository := profile.NewRepository(db)
	profileService := profile.NewService(profileRepository)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)

	app := NewApp()

	err = wails.Run(&options.App{
		Title:  "Dolphin | Finance Manager",
		Width:  1024,
		Height: 768,
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
			earningService,
			expenseService,
			profileService,
			categoryService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
