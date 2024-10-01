package main

import (
	"chipa.me/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, h handlers.Handlers) {
	e.Static("/static", "static")

	e.GET("/", h.Home())
}
