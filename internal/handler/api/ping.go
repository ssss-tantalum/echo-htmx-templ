package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "Pong!")
}
