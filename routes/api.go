package routes

import (
	"github.com/captainblue210/goliff/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/weather", api.FetchWeatherData())
	}
}
