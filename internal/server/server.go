package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/core"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/routes"
)

func RunServer() error {
	app := core.New()
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
	app.Echo.Static("/static", "static")

	routes.InitRoutes(app)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.SeverPort),
		Handler:      app.Echo,
		ReadTimeout:  time.Duration(app.Config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Config.WriteTimeout) * time.Second,
	}

	return server.ListenAndServe()
}
