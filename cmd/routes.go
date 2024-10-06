package main

import (
	"chipa.me/pkg/handlers"
	"chipa.me/pkg/md"
	"chipa.me/pkg/services"
	"chipa.me/views"
	"chipa.me/views/pages"
	"github.com/labstack/echo/v4"
)

func addRoutes(
	e *echo.Echo,
	h handlers.Handlers,
	postService *services.PostService,
	md *md.Markdown,
) {
	e.Static("/static", "static")

	e.GET("/", h.Home())

	e.GET("/blog", h.Blog(postService))
	e.GET("/blog/:slug", h.BlogEntry(postService, md))

	e.GET("/blog/edit", h.Edit())
	e.POST("/blog/edit", h.EditPost(md)) //HTMX

	// Test
	e.GET("/test", func(c echo.Context) error {
		return views.Render(c, 200, pages.Test())
	})
}
