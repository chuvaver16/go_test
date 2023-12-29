package provider

import "os"

type GeoLocation struct {
	City string
	Lat  float64
	Lon  float64
}

type Weather struct {
	Temp float64
}

type IWeatherProvider interface {
	GetCoordinate(city string) (*GeoLocation, error)
	GetWeatherByGeo(geo *GeoLocation) (*Weather, error)
	GetWeatherByCity(city string) (*Weather, error)
}

type WeatherProvider struct {
	UriGeo     string
	UriWeather string
	ApiKey     string
}

type IGeoProvider interface {
	GetCoordinate(ip string) (*GeoLocation, error)
}

type GeoProvider struct {
	UriGeo     string
	UriWeather string
	ApiKey     string
}

func InitWeatherProviders() map[string]IWeatherProvider {

	owm := new(OWM)
	owm.UriGeo = os.Getenv("OWM_URI_GEO")
	owm.UriWeather = os.Getenv("OWM_URI_WEATHER")
	owm.ApiKey = os.Getenv("OWM_API_KEY")

	om := new(OpenMeteo)
	om.UriGeo = os.Getenv("OPENMETEO_URI_GEO")
	om.UriWeather = os.Getenv("OPENMETEO_URI_WEATHER")

	providers := map[string]IWeatherProvider{
		"OpenMeteo": om,
		"OWM":       owm,
	}

	return providers
}

func InitGeoProvider() IGeoProvider {

	ipapi := new(IPApi)
	ipapi.UriGeo = os.Getenv("IPAPI_URI_GEO")

	return ipapi
}
