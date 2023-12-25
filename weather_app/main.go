// main.go
package main

import (
	"log"
	"weather_app/api"
	"weather_app/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	e := echo.New()

	api := api.New()

	e.Renderer = handler.InitTemplates()

	e.GET("/", handler.MainHandler)
	e.GET("/api/weather/", api.GetWeather)
	e.GET("/api/weather/:city", api.GetWeather)

	e.Start(":8081")
}
