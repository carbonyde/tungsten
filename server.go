package tungsten

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(app App) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	app.PageRouter(e)
	app.FunctionRouter(e)

	if Env.Watch {
		e.GET("/hot-reload", HotReload)
	}

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
