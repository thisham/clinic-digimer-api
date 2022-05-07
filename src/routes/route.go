package routes

import (
	"digimer-api/src/factories"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	// api := route.Group("api/")

	webFactory := factories.WebInit()
	// apiFactory := factories.ApiInit()

	// main
	route.GET("/", webFactory.InitialPageHandler.InitialPage)

	return route
}
