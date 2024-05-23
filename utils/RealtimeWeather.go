package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetRealtimeWeather(location string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("APIKEY")
	sampleURL := "https://api.tomorrow.io/v4/weather/realtime?location=%s&units=metric&apikey=%s"
	queryURL := fmt.Sprintf(sampleURL, location, apiKey)
	fmt.Println(queryURL)
}
