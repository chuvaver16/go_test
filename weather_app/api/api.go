package api

import (
	"errors"
	"log"
	"net/http"
	"weather_app/client"
	"weather_app/provider"

	"github.com/labstack/echo/v4"
)

type API struct {
	WeatherProviders map[string]provider.IWeatherProvider
	GeoProvider      provider.IGeoProvider
}

func New() *API {

	api := new(API)

	api.WeatherProviders = provider.InitWeatherProviders()
	api.GeoProvider = provider.InitGeoProvider()

	return api
}

func (api *API) GetWeather(c echo.Context) error {

	city := c.Param("city")

	log.Println(city)

	body := map[string]interface{}{
		"city": city,
	}

	if city == "" {
		ip := client.GetLocalIP()

		data, err := api.GeoProvider.GetCoordinate(ip)
		if err != nil {
			log.Print(err)
			body["error"] = err
			return c.JSON(http.StatusInternalServerError, body)
		} else if data == nil || data.City == "" {
			body["error"] = errors.New("City is not detected")
			return c.JSON(http.StatusInternalServerError, body)
		} else {
			city = data.City
			body["city"] = city
		}
	}

	for name, p := range api.WeatherProviders {
		data, err := p.GetWeatherByCity(city)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, body)
		}
		body[name] = data
	}

	return c.JSON(http.StatusOK, body)
}
