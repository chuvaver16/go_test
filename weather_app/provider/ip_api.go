package provider

import (
	"encoding/json"
	"log"
	"weather_app/client"
)

type IPApi struct {
	GeoProvider
}

type IPApiGeoResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func (ipapi *IPApi) GetCoordinate(ip string) (*GeoLocation, error) {

	var pparams = map[string]string{
		"{ip}": ip,
	}

	body, err := client.Get(ipapi.UriGeo, nil, pparams)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var data IPApiGeoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return nil, err
	}
	log.Print(data)
	geo := &GeoLocation{
		City: data.City,
		Lat:  data.Lat,
		Lon:  data.Lon,
	}

	return geo, nil
}
