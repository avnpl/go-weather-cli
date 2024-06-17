package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-weather-cli/utils"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	var multipleCities string
	var listAllData bool
	flag.StringVar(&multipleCities, "m", "", "List common weather data points for specified cities")
	flag.BoolVar(&listAllData, "a", false, "List all data points")
	flag.Parse()

	// If no arguments or flags are passed, open the interactive mode
	if len(os.Args) == 1 {
		DefaultAction()
	}

	// If only one argument is passed, print the common weather data
	if len(os.Args) == 2 {
		utils.PrintCommonWeatherData(os.Args[1])
		os.Exit(0)
	}

	// If 2 arguments are passed with one being `-a` then list all weather data of given city
	if len(os.Args) == 3 && listAllData {
		fmt.Println("List all data of given city")
		os.Exit(0)
	}

	// Split the string using spaces to get slice containing city names
	sliceOfCities := strings.Split(multipleCities, " ")
	fmt.Println(sliceOfCities)

	// If `-a` is passed, list all data of given cities or print common data points
	if listAllData {
		fmt.Println("Listing all data of given cities")
	} else {
		fmt.Println("Listing common data of given cities")
	}
}

func DefaultAction() {
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

		// Print the commonly used Weather datapoints
		utils.PrintCommonWeatherData(input)
	}
}
