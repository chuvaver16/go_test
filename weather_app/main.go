// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"weather_app/provider/owm"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	//e := echo.New()

	coordinates, err := owm.GetCoordinate("Moscow", "RU")
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(coordinates)

	lat := (*coordinates)[0].Lat
	lon := (*coordinates)[0].Lon

	weather, err := owm.GetWeather(lat, lon)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(weather)

	// Передача города в запросе
	//e.GET("/weather/:city", getWeather)

	// Старт сервера
	//e.Start(":8080")
}

func getWeather2(c echo.Context) error {
	city := c.Param("city")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"city": city,
	})
}
