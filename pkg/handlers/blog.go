package handlers

import (
	"net/http"
	"slices"

	"chipa.me/pkg/md"
	"chipa.me/pkg/services"
	"chipa.me/views"
	"chipa.me/views/pages"
	"github.com/labstack/echo/v4"
)

// Funcion que retorna una funcion
func (h Handlers) Blog(s *services.PostService) echo.HandlerFunc {
	return func(c echo.Context) error {

		posts := s.All()

		slices.SortFunc(posts, func(a *services.Post, b *services.Post) int {
			return b.CreatedAt.Compare(a.CreatedAt)
		})

		model := pages.BlogModel{
			BaseModel: pages.BaseModel{
				Title: "Blog",
			},
			Posts: posts,
		}

		return views.Render(c, http.StatusOK, pages.Blog(model))
	}
}

func (h Handlers) BlogEntry(s *services.PostService, md *md.Markdown) echo.HandlerFunc {
	return func(c echo.Context) error {
		slug := c.Param("slug")

		post, err := s.GetBySlug(slug)
		if err != nil {
			return c.HTML(404, "<h1>Post Not Found</h1>")
		}

		toc, err := md.RenderTOC(post.Content)
		if err != nil {
			h.Logger.Error().Err(err).Msg("[Blog] Error renderizando TOC")
		}
		body, err := md.RenderBody(post.Content)
		if err != nil {
			h.Logger.Error().Err(err).Msg("[Blog] Error renderizando Cuerpo")
		}

		model := pages.BlogEntryModel{
			BaseModel: pages.BaseModel{
				Title: "Home",
			},
			Post: body,
			Toc:  toc,
		}

		return views.Render(c, http.StatusOK, pages.BlogEntry(model))
	}
}
