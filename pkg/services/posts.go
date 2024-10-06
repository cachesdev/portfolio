package services

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/adrg/frontmatter"
	"golang.org/x/text/unicode/norm"
)

var (
	ErrPostDoesNotExist = errors.New("post does not exist")
)

type Post struct {
	Slug        string    `yaml:"slug"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	CreatedAt   time.Time `yaml:"CreatedAt"`
	UpdatedAt   time.Time `yaml:"UpdatedAt"`
	Tags        []string  `yaml:"tags"`
	Content     []byte    // unrendered
	HTMLContent string    // pre-rendered
}

type Slug = string
type Tags = string

type PostService struct {
	posts map[Slug]*Post
	mu    sync.RWMutex
}

// NewPostService returns an initialized *PostService already holding all posts
// and an error, which is composed of any errors during directory walking or parsing
func NewPostService() (*PostService, error) {
	posts, err := gatherPosts()
	return &PostService{
		posts: posts,
		mu:    sync.RWMutex{},
	}, err
}

func (s *PostService) GetBySlug(slug string) (*Post, error) {
	if slug == "" {
		return nil, ErrPostDoesNotExist
	}

	post, ok := s.posts[slug]
	if !ok {
		return nil, ErrPostDoesNotExist
	}

	return post, nil
}

func (s *PostService) All() []*Post {
	var posts []*Post
	for _, post := range s.posts {
		posts = append(posts, post)
	}

	return posts
}

// gatherPosts walks the root "posts" directory, reads any .md file and initializes
// a new Post struct for each of them.
func gatherPosts() (map[Slug]*Post, error) {
	posts := make(map[Slug]*Post)
	var errors []error

	err := filepath.WalkDir("posts",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				errors = append(errors, fmt.Errorf("error accessing %s: %w", path, err))
				return nil // Continue walking
			}

			if d.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}

			file, err := os.ReadFile(path)
			if err != nil {
				errors = append(errors, fmt.Errorf("error reading %s: %w", path, err))
				return nil // Continue to next file
			}

			post, err := parsePost(file)
			if err != nil {
				errors = append(errors, fmt.Errorf("error parsing %s: %w", path, err))
				return nil // Continue to next file
			}

			if post.Slug == "" {
				post.Slug = generateSlug(post.Title)
			}

			posts[post.Slug] = post
			return nil
		})

	if err != nil {
		errors = append(errors, fmt.Errorf("error walking directory: %w", err))
	}

	if len(errors) > 0 {
		return posts, fmt.Errorf("encountered %d errors during post gathering: %v", len(errors), errors)
	}

	return posts, nil
}

// parsePost trims the frontmatter and loads the data into the relevant
// struct fields, then returns the trimmed data.
func parsePost(content []byte) (*Post, error) {
	var post Post
	rest, err := frontmatter.Parse(bytes.NewReader(content), &post)
	if err != nil {
		return nil, err
	}

	if post.Slug == "" {
		post.Slug = generateSlug(post.Title)
	}

	post.Content = rest
	return &post, nil
}

// generateSlug returns a slug which will be composed based on the given title.
// if the title is empty, it will return a generated, default slug.
func generateSlug(title string) string {
	// Default in case of empty title
	if title == "" {
		return "untitled-post-" + time.Now().Format("20060102-150405")
	}

	// Normalize unicode (we separate letter and accent)
	title = norm.NFKD.String(title)

	var builder strings.Builder
	var lastWasInvalid bool

	// does not handle transliteration.
	for _, r := range title {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
			lastWasInvalid = false
		} else if !lastWasInvalid {
			builder.WriteRune('-')
			lastWasInvalid = true
		}
	}

	// Trim hyphens from start and end
	slug := strings.Trim(builder.String(), "-")

	return slug
}
