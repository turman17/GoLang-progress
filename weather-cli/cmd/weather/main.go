package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/turman17/weather-cli/internal/weather"
)

func main() {
	var city string
	var units string

	flag.StringVar(&city, "city", "", "City (e.g. Lisbon, Kyiv, Cairo)")
	flag.StringVar(&units, "units", "metric", "Units: standard | metric | imperial")
	flag.Parse()

	if city == "" {
		fmt.Fprintf(os.Stderr, "usage: %s -city <name> [-units metric|imperial|standard]\n", os.Args[0])
		os.Exit(2)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading .env file: %v\n", err)
		os.Exit(1)
	}

	apiKey := os.Getenv("OW_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "error: set OW_API_KEY environment variable")
		os.Exit(1)
	}

	client := weather.Client{
		APIKey:  apiKey,
		Timeout: 7 * time.Second,
		BaseURL: "https://api.openweathermap.org/data/2.5",
	}

	data, err := client.GetWeatherByCity(city, units)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(weather.FormatCurrent(data, units))
}