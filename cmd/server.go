package main

import (
	"chipa.me/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func newServer(
	logger zerolog.Logger,
	h handlers.Handlers,
) *echo.Echo {
	e := echo.New()

	addRoutes(e, h)

	return e
}
