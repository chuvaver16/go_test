package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"weather_app/client"
)

type OWM struct {
	WeatherProvider
}

type OWMGeoResponse []struct {
	City       string      `json:"name"`
	LocalNames interface{} `json:"local_names"`
	Country    string      `json:"country"`
	State      string      `json:"state"`
	Lat        float64     `json:"lat"`
	Lon        float64     `json:"lon"`
}

type OWMWeatherResponse struct {
	Coord   interface{} `json:"coord"`
	Weather interface{} `json:"weather"`
	Base    string      `json:"base"`
	Main    struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       interface{} `json:"wind"`
	Rain       interface{} `json:"rain"`
	Clouds     interface{} `json:"clouds"`
	Dt         int         `json:"dt"`
	Sys        interface{} `json:"sys"`
	Timezone   int         `json:"timezone"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
}

func (owm *OWM) GetCoordinate(city string) (*GeoLocation, error) {

	var qparams = map[string]string{
		"q":     city,
		"appid": owm.ApiKey,
	}

	body, err := client.Get(owm.UriGeo, qparams, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var data OWMGeoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	geo := &GeoLocation{
		City: data[0].City,
		Lat:  data[0].Lat,
		Lon:  data[0].Lon,
	}

	return geo, nil
}

func (owm *OWM) GetWeatherByGeo(geo *GeoLocation) (*Weather, error) {

	var qparams = map[string]string{
		"lat":   fmt.Sprintf("%f", geo.Lat),
		"lon":   fmt.Sprintf("%f", geo.Lon),
		"appid": owm.ApiKey,
		"units": "metric",
	}

	body, err := client.Get(owm.UriWeather, qparams, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var data OWMWeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	weather := &Weather{
		Temp: data.Main.Temp,
	}
	return weather, nil
}

func (owm *OWM) GetWeatherByCity(city string) (*Weather, error) {

	geo, err := owm.GetCoordinate(city)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	weather, err := owm.GetWeatherByGeo(geo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return weather, nil
}
