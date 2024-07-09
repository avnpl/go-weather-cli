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
func FetchCoOrdinates(location string) (float64, float64) {
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
		Handles JSON response as it can be either a single object or an array of objects since there can be multiple locations around the world which share the same name, in which case, the API returns an array of location data. The first one is assumed to be correct and used in this case

		- Initialize a variable of type `GeocodeAPIResp` to hold the JSON data
		- Attempt to unmarshal `resBody` into `resData`.
		- If the first unmarshal attempt fails, declare a slice `resDataArr` of `GeocodeAPIResp`.
		- Try to unmarshal `resBody` into `resDataArr`.
		- If both unmarshal attempts fail, log a fatal error, print the response body and error message.
		- If unmarshaling into `resDataArr` is successful, set `resData` to the first element of `resDataArr`.
	*/

	var resData GeocodeAPIResp
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		var resDataArr []GeocodeAPIResp
		newErr := json.Unmarshal(resBody, &resDataArr)
		if newErr != nil {
			log.Fatalf("Error converting GeoCoding response body to JSON struct...")
			fmt.Print(string(resBody))
			fmt.Println("Error Message:", err.Error())
		}
		resData = resDataArr[0]
	}

	/*
		Convert the latitude and longitude to float values
	*/
	lat, err := strconv.ParseFloat(resData.Lat, 32)
	if err != nil {
		fmt.Println(string(resBody))
		fmt.Println(resData.Lat)
		fmt.Println(location, err)
		log.Fatalf("Error converting Latitude to float64")
	}
	lon, err := strconv.ParseFloat(resData.Lon, 64)
	if err != nil {
		log.Fatalf("Error converting Longitude to float64...")
		fmt.Println(resData.Lon)
		fmt.Println(location, err)
	}

	return lat, lon
}
