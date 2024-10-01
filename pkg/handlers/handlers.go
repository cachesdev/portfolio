package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"go.abhg.dev/goldmark/toc"

	"chipa.me/pkg/services"
	"chipa.me/views"
	"chipa.me/views/pages"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Handlers struct {
	// Logger is the only global dependency of handlers.
	// the rest of the deps are passed as arguments
	Logger zerolog.Logger
}

func NewHandlers(logger zerolog.Logger) Handlers {
	return Handlers{
		Logger: logger,
	}
}

func (h Handlers) Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := services.GetPosts()
		if err != nil {
			fmt.Println(err)
		}

		markdown := goldmark.New(
			goldmark.WithExtensions(highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
				),
			)),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
		)

		bufTOC := bytes.Buffer{}
		doc := markdown.Parser().Parse(text.NewReader(posts[0]))
		tree, _ := toc.Inspect(doc, posts[0], toc.Compact(true))
		toc := toc.RenderList(tree)
		markdown.Renderer().Render(&bufTOC, posts[0], toc)

		bufBody := bytes.Buffer{}
		err = markdown.Convert(posts[0], &bufBody)
		if err != nil {
			h.Logger.Error().Err(err).Msg("No se pudo convertir blog")
		}

		model := pages.HomeModel{
			BaseModel: pages.BaseModel{
				Title: "Home",
			},
			Post: bufBody.String(),
			Toc:  bufTOC.String(),
		}

		return views.Render(c, http.StatusOK, pages.Home(model))
	}
}
