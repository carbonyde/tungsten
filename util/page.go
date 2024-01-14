package util

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Page struct {
	Method     string
	Route      string
	Component  func(c echo.Context) templ.Component
	middleware []echo.MiddlewareFunc
}

type Function struct {
	Method     string
	Route      string
	Handler    echo.HandlerFunc
	middleware []echo.MiddlewareFunc
}

type App struct {
	pages     []*Page
	functions []*Function
}

func (a *App) AddPage(pages ...*Page) {
	a.pages = append(a.pages, pages...)
}

func (a *App) AddFunction(functions ...*Function) {
	a.functions = append(a.functions, functions...)
}

func (a *App) PageRouter(e *echo.Echo) {
	for _, page := range a.pages {
		e.GET(page.Route, func(c echo.Context) error {
			return Render(c, page.Component(c))
		}, page.middleware...)
	}
}

func (a *App) FunctionRouter(e *echo.Echo) {
	for _, function := range a.functions {
		e.Add(function.Method, function.Route, function.Handler, function.middleware...)
	}
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
