package main

import (
	"bufio"
	"fmt"
	"go-weather-cli/utils"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Create a buffered reader to read the input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type 'x' at any moment to exit")
	fmt.Println()
	for {
		fmt.Println("> Enter the city name : ")
		fmt.Print("> ")

		/*
			- Read the input from user into a variable
			- Log the error, if any, and exit
		*/
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Error while reading the input..")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		/*
			- Trim whitespaces from the input
			- PathEscape the string so that it can be included in the URL
		*/
		input = strings.TrimSpace(input)
		input = url.PathEscape(input)

		// Exit if the user enters 'x'
		if input == "x" {
			fmt.Print("Exiting...")
			os.Exit(0)
		}

		// Get the latitude and longitude of the place
		lat, lon := utils.GeoCodingAPIClient(input)

		// Fetch data from the API
		apiData := utils.WeatherAPIClient(lat, lon)

		// Print the commonly used Weather datapoints
		utils.PrintCommonWeatherData(apiData, lat, lon)
	}
}
