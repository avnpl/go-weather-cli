package utils

import (
	"fmt"
)

func GetCommonWeatherData(location string) CommonWeatherData {
	apiData := WeatherAPIClient(location)

	return CommonWeatherData{
		CloudCover:          apiData.Data.Values.CloudCover,
		Humidity:            apiData.Data.Values.Humidity,
		RainIntensity:       apiData.Data.Values.RainIntensity,
		SnowIntensity:       apiData.Data.Values.SnowIntensity,
		Temperature:         apiData.Data.Values.Temperature,
		TemperatureApparent: apiData.Data.Values.TemperatureApparent,
		Visibility:          apiData.Data.Values.Visibility,
		WindDirection:       apiData.Data.Values.WindDirection,
		WindSpeed:           apiData.Data.Values.WindSpeed,
		Location:            apiData.Location.Name,
	}
}

func PrintCommonWeatherData(location string) {
	data := GetCommonWeatherData(location)

	fmt.Println("----------------------------------")
	fmt.Printf("Temperature is          %.1f °C\n", data.Temperature)
	fmt.Printf("And Feels like          %.1f °C\n", data.TemperatureApparent)
	fmt.Printf("Cloud Cover is          %.1f %%\n", data.CloudCover)
	fmt.Printf("Humidity is             %.1f %%\n", data.Humidity)
	fmt.Printf("Possibility of Rain is  %.1f %%\n", data.RainIntensity)
	fmt.Printf("Possibility of Snow is  %.1f %%\n", data.SnowIntensity)
	fmt.Printf("Visibility is           %.1f km\n", data.Visibility)
	fmt.Printf("Wind Direction is       %.1f degrees\n", data.WindDirection)
	fmt.Printf("Wind Speed is           %.1f m/s\n", data.WindSpeed)
	fmt.Println("----------------------------------")
}
