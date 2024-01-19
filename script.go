package tungsten

import (
	"github.com/a-h/templ"
)

func Script(source string, variables map[string]string) templ.Component {
	return templ.Raw("<script async type='module'>" + Build(source, variables) + "</script>")
}

func InlineScript(source string) templ.Component {
	return templ.Raw("<script async type='module'>" + BuildInline(source) + "</script>")
}
