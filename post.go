package tungsten

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Post struct {
	Name        string
	Path        string
	Title       string
	Description string
	Tags        []string
	Keywords    []string
	PublishedAt string
	ReadingTime int
}

type FullPost struct {
	Name        string
	Path        string
	Title       string
	Description string
	Tags        []string
	Keywords    []string
	PublishedAt string
	ReadingTime int
	Prev        Post
	Next        Post
}

type Frontmatter struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Keywords    []string `yaml:"keywords"`
	PublishedAt string   `yaml:"publishedAt"`
}

func GetPosts(path string) []FullPost {
	entries, err := os.ReadDir("./content")

	if err != nil {
		panic(err)
	}

	posts := []Post{}

	for _, e := range entries {
		post, err := os.ReadFile(filepath.Join("content", e.Name()))

		if err != nil {
			panic(err)
		}

		meta, _, err := ConvertMarkdown(post)

		if err != nil {
			panic(err)
		}

		posts = append(posts, Post{
			Name:        strings.Split(e.Name(), ".")[0],
			Path:        path,
			Title:       meta.Title,
			Description: meta.Description,
			Keywords:    meta.Keywords,
			Tags:        meta.Tags,
			PublishedAt: meta.PublishedAt,
			ReadingTime: len(string(post[:])) / 285,
		})

	}

	sort.Slice(posts, func(i, j int) bool {
		a, err := time.Parse(time.DateOnly, posts[i].PublishedAt)

		if err != nil {
			panic(err)
		}
		b, err := time.Parse(time.DateOnly, posts[j].PublishedAt)

		if err != nil {
			panic(err)
		}

		return b.Before(a)
	})

	fullPosts := []FullPost{}

	for index, post := range posts {
		var prev = Post{}
		var next = Post{}

		if index < len(posts)-1 {
			prev = posts[index+1]
		}
		if index > 0 {
			next = posts[index-1]
		}

		fullPosts = append(fullPosts, FullPost{
			Name:        post.Name,
			Path:        post.Path,
			Title:       post.Title,
			Description: post.Description,
			Keywords:    post.Keywords,
			Tags:        post.Tags,
			PublishedAt: post.PublishedAt,
			ReadingTime: post.ReadingTime,
			Prev:        prev,
			Next:        next,
		})
	}

	return fullPosts
}
