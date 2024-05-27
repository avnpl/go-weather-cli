package utils

import (
	"fmt"
	"reflect"
)

func GetCommonWeatherData(apiData APIResp) CommonWeatherData {
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

func PrintCommonWeatherData(apiData APIResp) {
	data := GetCommonWeatherData(apiData)
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fmt.Printf("%s: %v\n", fieldName, field)
	}
}
