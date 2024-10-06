package md

import (
	"bytes"
	"fmt"

	chromastyles "github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/toc"
)

type Markdown struct {
	goldmark.Markdown
}

// NewMarkdown returns a wrapper for goldmark.Markdown, and provides some additional
// helper functions, while also doing some basic configuration.
func NewMarkdown() *Markdown {
	hlExt := highlighting.NewHighlighting(
		highlighting.WithCustomStyle(chromastyles.Get("gruvbox")),
	)

	markdown := goldmark.New(
		goldmark.WithExtensions(hlExt),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	return &Markdown{
		Markdown: markdown,
	}
}

func (md *Markdown) RenderTOC(post []byte) (string, error) {
	bufTOC := bytes.Buffer{}
	doc := md.Parser().Parse(text.NewReader(post))

	tree, err := toc.Inspect(doc, post, toc.Compact(true))
	if err != nil {
		return "", fmt.Errorf("[RenderTOC]: %w", err)
	}

	toc := toc.RenderList(tree)
	err = md.Renderer().Render(&bufTOC, post, toc)
	if err != nil {
		return "", fmt.Errorf("[RenderTOC]: %w", err)
	}

	return bufTOC.String(), nil
}

func (md *Markdown) RenderBody(post []byte) (string, error) {
	bufBody := bytes.Buffer{}
	err := md.Convert(post, &bufBody)
	if err != nil {
		return "", fmt.Errorf("[RenderBody]: %w", err)
	}

	return bufBody.String(), nil
}
