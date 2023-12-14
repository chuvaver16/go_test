package owm

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"weather_app/client"
	"weather_app/dto"
)

func GetCoordinate(city string, country string) (*dto.OWMGeoResponse, error) {

	uri := os.Getenv("OWM_URI_GEO")
	key := os.Getenv("OWM_API_KEY")

	var qparams = map[string]string{
		"q":     city + "," + country,
		"appid": key,
	}

	body, err := client.Get(uri, qparams)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	data := new(dto.OWMGeoResponse)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	return data, nil
}

func GetWeather(lat float64, lon float64) (*dto.OWMWeatherResponse, error) {

	uri := os.Getenv("OWM_URI_WEATHER")
	key := os.Getenv("OWM_API_KEY")

	var qparams = map[string]string{
		"lat":   fmt.Sprintf("%f", lat),
		"lon":   fmt.Sprintf("%f", lon),
		"appid": key,
		"units": "metric",
	}

	body, err := client.Get(uri, qparams)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	data := new(dto.OWMWeatherResponse)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	return data, nil
}
