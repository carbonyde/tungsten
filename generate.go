package tungsten

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Route struct {
	Name string
	Path string
}

func Generate(routes []Route, contentRoute string, contentDir string) error {
	posts, err := os.ReadDir(contentDir)

	if err != nil {
		return err
	}

	for _, post := range posts {
		name := strings.Split(post.Name(), ".")[0]
		routes = append(routes, Route{
			Name: name,
			Path: "/" + contentRoute + "/" + name,
		})
	}

	for _, route := range routes {
		response, err := http.Get("http://localhost:8080" + route.Path)

		if err != nil {
			return err
		}

		body, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join("dist", route.Path), os.ModePerm); err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join("dist", route.Path, "index.html"), body, 0666); err != nil {
			return err
		}
	}

	return nil
}
