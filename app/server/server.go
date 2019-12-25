package server

import (
	"fmt"
	"go-id-bot/app/server/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	httpServer *http.Server

	Version string

	Secret string
}

func (s *Server) Run(host string, port int) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", api.Ping)
	e.POST(fmt.Sprintf("/bot:%s", s.Secret), api.Bot)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
