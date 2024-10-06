package main

import (
	"chipa.me/pkg/handlers"
	"chipa.me/pkg/md"
	"chipa.me/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func newServer(
	logger zerolog.Logger,
	h handlers.Handlers,
	postService *services.PostService,
	md *md.Markdown,
) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	// 	Level: 6,
	// }))

	addRoutes(e, h, postService, md)

	return e
}
