package ui

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

func LucideIcon(name string, class string) templ.Component {
	icon, err := os.ReadFile(filepath.Join("./node_modules/lucide-static/icons", name+".svg"))

	if err != nil {
		println("Error: to use LucideIcon install lucide-static from NPM first")

		panic(err)
	}

	svg := strings.Replace(string(icon[:]), "<svg", "<svg class='"+class+"'", 1)

	return templ.Raw(svg)
}
