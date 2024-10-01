package services

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type post = []byte

type posts []post

func GetPosts() (posts, error) {
	var data posts
	err := filepath.WalkDir("posts",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("WalkDir: %w", err)
			}

			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".md" {
				return nil
			}

			file, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("ReadFile: %w", err)
			}

			data = append(data, file)

			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("getPosts: %w", err)
	}

	return data, nil
}

type TocItem struct {
	ID       string
	Title    string
	Level    int
	Children []*TocItem
}

func GenerateTOC(markdown []byte) ([]*TocItem, error) {
	doc := goldmark.DefaultParser().Parse(text.NewReader(markdown))
	tocItems := []*TocItem{}
	stack := []*TocItem{}

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if n.Kind() == ast.KindHeading {
			heading := n.(*ast.Heading)
			title := string(n.Text(markdown))
			id := strings.ToLower(strings.ReplaceAll(title, " ", "-"))

			item := &TocItem{
				ID:    id,
				Title: title,
				Level: heading.Level,
			}

			for len(stack) > 0 && stack[len(stack)-1].Level >= item.Level {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, item)
			} else {
				tocItems = append(tocItems, item)
			}

			stack = append(stack, item)
		}

		return ast.WalkContinue, nil
	})

	return tocItems, nil
}

func RenderTOC(items []*TocItem) string {
	var buf bytes.Buffer
	buf.WriteString("<ul class='toc'>")
	for _, item := range items {
		renderTOCItem(&buf, item)
	}
	buf.WriteString("</ul>")
	return buf.String()
}

func renderTOCItem(buf *bytes.Buffer, item *TocItem) {
	buf.WriteString(fmt.Sprintf("<li><a href='#%s'>%s</a>", item.ID, item.Title))
	if len(item.Children) > 0 {
		buf.WriteString("<ul>")
		for _, child := range item.Children {
			renderTOCItem(buf, child)
		}
		buf.WriteString("</ul>")
	}
	buf.WriteString("</li>")
}
