package core

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type App struct {
	Echo   *echo.Echo
	Config *AppConfig

	DB *bun.DB
}

func New() *App {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	app := &App{
		Echo:   echo.New(),
		Config: config,
		DB:     db,
	}

	return app
}
