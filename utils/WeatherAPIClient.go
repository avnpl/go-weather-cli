package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

/*
Take the `location` name as the input string and return the API response as a struct
*/
func WeatherAPIClient(location string) WeatherAPIResp {
	// Loading the Environment Variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading the env variables..")
		fmt.Println(err.Error())
	}

	// Getting the APIKEY from .env file
	apiKey := os.Getenv("WEATHERAPIKEY")

	// Get the co-ordinates of the location
	lat, lon := GeoCodingAPIClient(location)

	/*
		- Convert latitude and longitude to string to be used in the URL
		- PathEscape the string so that it can be included in the URL
	*/
	location = fmt.Sprintf("%f,%f", lat, lon)
	location = url.PathEscape(location)

	/*
		- Create the query URL
		- Send a GET request
		- Log the error, if any
	*/
	queryURL := fmt.Sprintf("https://api.tomorrow.io/v4/weather/realtime?location=%s&units=metric&apikey=%s", location, apiKey)
	res, err := http.Get(queryURL)
	if err != nil {
		log.Fatalf("Error fetching from API...")
		fmt.Println(err.Error())
	}

	// Read the response body into an array of bytes and log the error, if any
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body...")
		fmt.Println(err.Error())
	}

	// Defer the closing of the reader before the function returns
	defer res.Body.Close()

	/*
		- Create a variable to hold the JSON data using the type struct
		- Unmarshal the byte array data into JSON and store it in the `resData` variable
		- Throw error, if any
	*/
	var resData WeatherAPIResp
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		log.Fatalf("Error converting response body to JSON struct...")
		fmt.Println(err.Error())
	}

	// Returns the response data in JSON format
	return resData
}

/*
Print all the fields in the struct with its values
*/
func PrintAllVals(apiData WeatherAPIResp) {
	val := reflect.ValueOf(apiData.Data.Values)
	typ := reflect.TypeOf(apiData.Data.Values)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fmt.Printf("%s: %v\n", fieldName, field)
	}
	fmt.Print(apiData.Location.Name)
}
