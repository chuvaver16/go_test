package api

import (
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

	body := map[string]interface{}{}
	body["city"] = city

	if city == "" {
		ip := client.GetLocalIP()

		data, err := api.GeoProvider.GetCoordinate(ip)
		if err != nil {
			log.Print(err)
		}
		if data == nil || data.City == "" {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status": "error",
			})
		} else {
			city = data.City
		}

		//log.Print(data)
	}

	for name, p := range api.WeatherProviders {
		data, err := p.GetWeatherByCity(city)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status": "error",
			})
		}
		body[name] = data
	}

	return c.JSON(http.StatusOK, body)
}
