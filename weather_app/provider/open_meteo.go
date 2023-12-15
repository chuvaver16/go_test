package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"weather_app/client"
)

type OpenMeteo struct {
	WeatherProvider
}

type OpenMeteoGeoResponse struct {
	Results []struct {
		ID          int      `json:"id"`
		Name        string   `json:"name"`
		Latitude    float64  `json:"latitude"`
		Longitude   float64  `json:"longitude"`
		Elevation   float64  `json:"elevation"`
		FeatureCode string   `json:"feature_code"`
		CountryCode string   `json:"country_code"`
		Admin1ID    int      `json:"admin1_id"`
		Timezone    string   `json:"timezone"`
		Population  int      `json:"population,omitempty"`
		CountryID   int      `json:"country_id"`
		Country     string   `json:"country"`
		Admin1      string   `json:"admin1"`
		Admin2ID    int      `json:"admin2_id,omitempty"`
		Postcodes   []string `json:"postcodes,omitempty"`
		Admin2      string   `json:"admin2,omitempty"`
		Admin3ID    int      `json:"admin3_id,omitempty"`
		Admin3      string   `json:"admin3,omitempty"`
	} `json:"results"`
	GenerationtimeMs float64 `json:"generationtime_ms"`
}

type OpenMeteoWeatherResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		Time          string `json:"time"`
		Interval      string `json:"interval"`
		Temperature2M string `json:"temperature_2m"`
	} `json:"current_units"`
	Current struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Temperature2M float64 `json:"temperature_2m"`
	} `json:"current"`
}

func (om *OpenMeteo) GetCoordinate(city string) (*GeoLocation, error) {

	var qparams = map[string]string{
		"name": city,
	}

	body, err := client.Get(om.UriGeo, qparams, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var data OpenMeteoGeoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	if data.Results != nil {
		geo := &GeoLocation{
			City: data.Results[0].Name,
			Lat:  data.Results[0].Latitude,
			Lon:  data.Results[0].Longitude,
		}

		return geo, nil
	}

	return nil, errors.New("Service is wrong")

}

func (om *OpenMeteo) GetWeatherByGeo(geo *GeoLocation) (*Weather, error) {

	var qparams = map[string]string{
		"latitude":  fmt.Sprintf("%f", geo.Lat),
		"longitude": fmt.Sprintf("%f", geo.Lon),
		"current":   "temperature_2m",
	}

	body, err := client.Get(om.UriWeather, qparams, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var data OpenMeteoWeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}

	weather := &Weather{
		Temp: data.Current.Temperature2M,
	}
	return weather, nil

}

func (om *OpenMeteo) GetWeatherByCity(city string) (*Weather, error) {

	geo, err := om.GetCoordinate(city)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	weather, err := om.GetWeatherByGeo(geo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return weather, nil
}
