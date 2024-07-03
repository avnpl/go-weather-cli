package utils

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

func FetchCommonDataForCities(cities []string) []CommonWeatherData {
	var wg sync.WaitGroup
	WeatherData := make([]CommonWeatherData, 0, len(cities))
	// Create a channel which can store the data of the cities
	dataChannel := make(chan CommonWeatherData, len(cities))

	// Create a series of GoRoutines for each city and push them into the channel
	for i := 0; i < len(cities); i++ {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			data := FetchCommonWeatherDetails(city)
			dataChannel <- data
		}(cities[i])
	}

	// Wait for all GoRoutines to finish
	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	// Get the data from the channel and append it into the slice
	for data := range dataChannel {
		WeatherData = append(WeatherData, data)
	}

	// Return the slice
	return WeatherData
}

func PrintCommonWeatherDataCities(cities []string, data []CommonWeatherData) {
	typ := reflect.TypeOf(data[0])
	numOfFields := typ.NumField()
	numOfCities := len(cities)
	commonDataUnits := [...]string{"%", "%", "%", "%", "°C", "°C", "km", "°", "m/s"}

	// Make a slice containing the field names
	fieldNames := make([]string, numOfFields)
	for i := 0; i < numOfFields; i++ {
		fieldNames[i] = typ.Field(i).Name
	}

	// Find the maximum length of a field name
	var maxlen int
	for _, field := range fieldNames {
		if len(field) > maxlen {
			maxlen = len(field)
		}
	}

	// Print the header of the table
	fmt.Println()
	fmt.Print(strings.Repeat(" ", maxlen+2))
	for i := 0; i < numOfCities; i++ {
		fmt.Print("| ")
		fmt.Print(cities[i])
		fmt.Print(" ")
	}
	fmt.Print("|")
	fmt.Println()

	// Calculate the length of the header and print the separator
	totalHeaderLen := maxlen
	for _, city := range cities {
		if len(city) < 5 {
			totalHeaderLen += 7
		} else {
			totalHeaderLen += len(city) + 2
		}
	}
	totalHeaderLen += numOfCities + 3 + 4
	fmt.Println(strings.Repeat("-", totalHeaderLen))

	// Print fields in each cities one by one
	for i := 0; i < numOfFields; i++ {
		field := typ.Field(i).Name
		fmt.Print(field)
		fmt.Print(strings.Repeat(" ", maxlen+2-len(field)))
		fmt.Print("|")
		for j := 0; j < numOfCities; j++ {
			columnWidth := len(cities[j]) + 2
			if len(cities[j]) < 5 {
				columnWidth = 7
			}

			val := reflect.ValueOf(data[j])
			fieldVal := val.Field(i).Interface()
			valString := fmt.Sprintf("%v", fieldVal)

			initialSpace := (columnWidth - len(valString)) / 2

			fmt.Print(strings.Repeat(" ", initialSpace))
			fmt.Print(valString)
			fmt.Print(strings.Repeat(" ", columnWidth-initialSpace-len(valString)))
			fmt.Print("|")
		}
		fmt.Print(" ")
		fmt.Print(commonDataUnits[i])
		fmt.Println()
	}
	fmt.Println()
}

func FetchAllDataForCities(cities []string) []WeatherAPIResValues {
	var wg sync.WaitGroup
	WeatherData := make([]WeatherAPIResValues, 0, len(cities))
	// Create a channel which can store the data of the cities
	dataChannel := make(chan WeatherAPIResValues, len(cities))

	// Create a series of GoRoutines for each city and push them into the channel
	for i := 0; i < len(cities); i++ {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			data := FetchAllWeatherDetails(city)
			dataChannel <- data
		}(cities[i])
	}

	// Wait for all GoRoutines to finish
	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	// Get the data from the channel and append it into the slice
	for data := range dataChannel {
		WeatherData = append(WeatherData, data)
	}

	// Return the slice
	return WeatherData
}

func PrintAllWeatherDataForCities(cities []string, data []WeatherAPIResValues) {
	typ := reflect.TypeOf(data[0])
	numOfFields := typ.NumField()
	numOfCities := len(cities)

	// Make a slice containing the field names
	fieldNames := make([]string, numOfFields)
	for i := 0; i < numOfFields; i++ {
		fieldNames[i] = typ.Field(i).Name
	}

	// Find the maximum length of a field name
	var maxlen int
	for _, field := range fieldNames {
		if len(field) > maxlen {
			maxlen = len(field)
		}
	}

	// Print the header of the table
	fmt.Println()
	fmt.Print(strings.Repeat(" ", maxlen+2))
	for i := 0; i < numOfCities; i++ {
		fmt.Print("| ")
		fmt.Print(cities[i])
		fmt.Print(" ")
	}
	fmt.Print("|")
	fmt.Println()

	// Calculate the length of the header and print the separator
	totalHeaderLen := maxlen
	for _, city := range cities {
		if len(city) < 5 {
			totalHeaderLen += 7
		} else {
			totalHeaderLen += len(city) + 2
		}
	}
	totalHeaderLen += numOfCities + 3
	fmt.Println(strings.Repeat("-", totalHeaderLen))

	// Print fields in each cities one by one
	for i := 0; i < numOfFields; i++ {
		field := typ.Field(i).Name
		fmt.Print(field)
		fmt.Print(strings.Repeat(" ", maxlen+2-len(field)))
		fmt.Print("|")
		for j := 0; j < numOfCities; j++ {
			columnWidth := len(cities[j]) + 2
			if len(cities[j]) < 5 {
				columnWidth = 7
			}

			val := reflect.ValueOf(data[j])
			fieldVal := val.Field(i).Interface()
			valString := fmt.Sprintf("%v", fieldVal)

			initialSpace := (columnWidth - len(valString)) / 2

			fmt.Print(strings.Repeat(" ", initialSpace))
			fmt.Print(valString)
			fmt.Print(strings.Repeat(" ", columnWidth-initialSpace-len(valString)))
			fmt.Print("|")
		}
		fmt.Println()
	}
	fmt.Println()
}
