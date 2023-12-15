// main.go
package main

import (
	"log"
	"net/http"
	"weather_app/provider"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var providers map[string]provider.IProvider

func main() {
	providers = provider.InitProviders()

	e := echo.New()

	e.GET("/weather/:city", getWeather)

	e.Start(":8081")
}

func getWeather(c echo.Context) error {
	city := c.Param("city")

	body := map[string]interface{}{}
	body["city"] = city

	for name, p := range providers {
		data, err := p.GetWeatherByCity(city)
		if err != nil {
			log.Print(err)
		}
		body[name] = data
	}

	return c.JSON(http.StatusOK, body)
}
