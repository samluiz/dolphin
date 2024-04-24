package main

import (
	"embed"
	"fintrack/backend/db"
	"fintrack/backend/earning"
	"fintrack/backend/expense"
	"fintrack/backend/expense/category"
	"fintrack/backend/profile"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	db, err := db.Connect()

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
		Title:  "Fintrack | Finance Tracker",
		Width:  1024,
		Height: 768,
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
