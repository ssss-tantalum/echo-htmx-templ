package routes

import (
	"github.com/ssss-tantalum/echo-templ-htmx/internal/core"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/handler"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/handler/api"
)

func InitRoutes(app *core.App) {
	// HTMX Page Route
	indexHandler := handler.NewIndexHandler(app)
	app.Echo.GET("/", indexHandler.Index)

	// API Route
	pingHandler := api.NewPingHandler()
	app.Echo.GET("/api/ping", pingHandler.Pong)
}
