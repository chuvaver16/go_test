package dto

type OWMGeoResponse []struct {
	City       string      `json:"name"`
	LocalNames interface{} `json:"local_names"`
	Country    string      `json:"country"`
	State      string      `json:"state"`
	Lat        float64     `json:"lat"`
	Lon        float64     `json:"lon"`
}

type OWMWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
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
