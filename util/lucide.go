package util

import (
	"os"
	"path/filepath"
	"strings"
)

func LucideIcon(name string, class string) string {
	icon, err := os.ReadFile(filepath.Join("./node_modules/lucide-static/icons", name+".svg"))

	if err != nil {
		panic(err)
	}

	return strings.Replace(string(icon[:]), "<svg", "<svg class='"+class+"'", 1)
}
