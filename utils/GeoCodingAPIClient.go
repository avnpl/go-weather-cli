package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*
Take the `location` name as the input string and return the latitude and longitude as strings
*/
func GeoCodingAPIClient(location string) (float64, float64) {
	// Loading the Environment Variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading the env variables..")
		fmt.Println(err.Error())
	}

	// Getting the APIKEY from .env file
	apiKey := os.Getenv("GEOCODEAPIKEY")

	/*
		- Create the query URL
		- Send a GET request
		- Log the error, if any
	*/
	queryURL := fmt.Sprintf("https://us1.locationiq.com/v1/search?q=%s&format=json&key=%s", location, apiKey)
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
	var resData GeocodeAPIResp
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		log.Fatalf("Error converting GeoCoding response body to JSON struct...")
		fmt.Println(resBody)
		fmt.Println(err.Error())
	}

	/*
		Convert the latitude and longitude to float values
	*/
	lat, err := strconv.ParseFloat(resData[0].Lat, 32)
	if err != nil {
		log.Fatalf("Error converting Latitude to float64...")
		fmt.Println(err.Error())
	}
	lon, err := strconv.ParseFloat(resData[0].Lon, 32)
	if err != nil {
		log.Fatalf("Error converting Latitude to float64...")
		fmt.Println(err.Error())
	}

	return lat, lon
}
