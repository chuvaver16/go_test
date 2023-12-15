// main.go
package main

import (
	"log"
	"net/http"
	"weather_app/client"
	"weather_app/provider"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var weather_providers map[string]provider.IWeatherProvider
var geo_provider provider.IGeoProvider

func main() {
	weather_providers = provider.InitWeatherProviders()
	geo_provider = provider.InitGeoProvider()

	e := echo.New()

	e.GET("/weather/:city", getWeather)
	e.GET("/weather", getWeather)

	e.Start(":8081")
}

func getWeather(c echo.Context) error {
	city := c.Param("city")

	body := map[string]interface{}{}
	body["city"] = city

	if city == "" {
		ip := client.GetLocalIP()

		data, err := geo_provider.GetCoordinate(ip)
		if err != nil {
			log.Print(err)
		}
		if data == nil || data.City == "" {
			return c.JSON(http.StatusInternalServerError, nil)
		} else {
			city = data.City
		}

		//log.Print(data)
	}

	for name, p := range weather_providers {
		data, err := p.GetWeatherByCity(city)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		body[name] = data
	}

	return c.JSON(http.StatusOK, body)
}
