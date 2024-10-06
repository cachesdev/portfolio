package handlers

import (
	"net/http"

	"chipa.me/pkg/md"
	"chipa.me/views"
	"chipa.me/views/pages"
	"github.com/labstack/echo/v4"
)

func (h Handlers) EditPost(md *md.Markdown) echo.HandlerFunc { // HTMX
	return func(c echo.Context) error {
		input := c.FormValue("content")

		body, err := md.RenderBody([]byte(input))
		if err != nil {
			h.Logger.Error().Err(err).Msg("[EditPost] Render error on Body")
			return c.HTML(http.StatusOK,
				`<div role="alert" class="alert alert-error">Markdown Rendering Error</div>`,
			)
		}

		return c.HTML(200, body)
	}
}

func (h Handlers) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {

		model := pages.EditorModel{
			BaseModel: pages.BaseModel{
				Title: "Editor Markdown",
			},
		}

		return views.Render(c, 200, pages.Editor(model))
	}
}
