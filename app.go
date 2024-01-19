package tungsten

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Page struct {
	Method      string
	Path        string
	Component   func(c echo.Context) templ.Component
	Middlewares []echo.MiddlewareFunc
}

type Function struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

type App struct {
	Pages     []Page
	Functions []Function
}

var app = App{}

func AddPage(pages Page) {
	app.Pages = append(app.Pages, pages)
}

func AddFunction(functions Function) {
	app.Functions = append(app.Functions, functions)
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func getMethod(method string) string {
	if len(method) > 0 {
		return method
	}
	return http.MethodGet
}

func pageHandler(e *echo.Echo, page Page) {
	e.Add(getMethod(page.Method), page.Path, func(c echo.Context) error {
		return Render(c, page.Component(c))
	}, page.Middlewares...)
}

func functionHandler(e *echo.Echo, function Function) {
	e.Add(getMethod(function.Method), function.Path, function.Handler, function.Middlewares...)
}

type Config struct {
	Address string
}

func getConfig(config Config) Config {
	if len(config.Address) == 0 {
		config.Address = "0.0.0.0:8080"
	}
	return config
}

func Start(c Config) {
	config := getConfig(c)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	for _, page := range app.Pages {
		pageHandler(e, page)
	}

	for _, function := range app.Functions {
		functionHandler(e, function)
	}

	if Env.Watch {
		e.GET("/hot-reload", HotReload)
	}

	e.Logger.Fatal(e.Start(config.Address))
}
