package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func Bot(c echo.Context) error {
	return c.String(http.StatusOK, "")
}