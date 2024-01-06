package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/core"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/handler/utils"
	"github.com/ssss-tantalum/echo-templ-htmx/templates/pages"
)

type IndexHandler struct {
	app *core.App
}

func NewIndexHandler(app *core.App) *IndexHandler {
	return &IndexHandler{
		app: app,
	}
}

func (h *IndexHandler) Index(c echo.Context) error {
	return utils.Render(c, pages.Index("tantalum!"))
}
