package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
)

func FetchWeatherData() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := os.Getenv("OPEN_WEATHER_API_KEY")
		url := "https://api.openweathermap.org/data/2.5/weather?id=1850144&appid=" + key

		cd, body, err := fasthttp.Get(nil, url)
		if err != nil && cd != fasthttp.StatusOK {
			logrus.Fatalf("Error get weather data %v", err)
		}

		return c.JSON(fasthttp.StatusOK, string(body))
	}
}
