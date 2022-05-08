package routes

import (
	"digimer-api/src/factories"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	api := route.Group("/api")

	webFactory := factories.WebInit()
	apiFactory := factories.ApiInit()

	// main
	route.GET("/", webFactory.InitialPageHandler.InitialPage)

	// api
	// polyclinics
	api.GET("/polyclinics", apiFactory.Polyclinic.ShowAllPolyclinicsHandler)
	api.GET("/polyclinics/:id", apiFactory.Polyclinic.ShowPolyclinicByIDHandler)
	api.POST("/polyclinics", apiFactory.Polyclinic.CreatePolyclinicHandler)
	api.PUT("/polyclinics/:id", apiFactory.Polyclinic.AmendPolyclinicByIDHandler)
	api.DELETE("/polyclinics/:id", apiFactory.Polyclinic.RemovePolyclinicByIDHandler)

	return route
}
