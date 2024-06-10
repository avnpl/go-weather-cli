package utils

import "time"

type WeatherAPIResp struct {
	Data struct {
		Time   time.Time           `json:"time"`
		Values WeatherAPIResValues `json:"values"`
	} `json:"data"`
	Location WeatherAPIResLocation `json:"location"`
}

type WeatherAPIResLocation struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}

type WeatherAPIResValues struct {
	CloudBase                float64 `json:"cloudBase"`
	CloudCeiling             float64 `json:"cloudCeiling"`
	CloudCover               float64 `json:"cloudCover"`
	DewPoint                 float64 `json:"dewPoint"`
	FreezingRainIntensity    float64 `json:"freezingRainIntensity"`
	Humidity                 float64 `json:"humidity"`
	PrecipitationProbability float64 `json:"precipitationProbability"`
	PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
	RainIntensity            float64 `json:"rainIntensity"`
	SleetIntensity           float64 `json:"sleetIntensity"`
	SnowIntensity            float64 `json:"snowIntensity"`
	Temperature              float64 `json:"temperature"`
	TemperatureApparent      float64 `json:"temperatureApparent"`
	UvHealthConcern          float64 `json:"uvHealthConcern"`
	UvIndex                  float64 `json:"uvIndex"`
	Visibility               float64 `json:"visibility"`
	WeatherCode              float64 `json:"weatherCode"`
	WindDirection            float64 `json:"windDirection"`
	WindGust                 float64 `json:"windGust"`
	WindSpeed                float64 `json:"windSpeed"`
}

type CommonWeatherData struct {
	CloudCover          float64 `json:"cloudCover"`
	Humidity            float64 `json:"humidity"`
	RainIntensity       float64 `json:"rainIntensity"`
	SnowIntensity       float64 `json:"snowIntensity"`
	Temperature         float64 `json:"temperature"`
	TemperatureApparent float64 `json:"temperatureApparent"`
	Visibility          float64 `json:"visibility"`
	WindDirection       float64 `json:"windDirection"`
	WindSpeed           float64 `json:"windSpeed"`
	Location            string  `json:"name"`
}

type GeocodeAPIResp []struct {
	PlaceID     string   `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       string   `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon,omitempty"`
}
