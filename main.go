package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// created by Tejaswi Cheripally - 500229934
// Define a struct named Weather to represent weather-related data
type Weather struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	FeelsLike   string `json:"feelslike"`
	Weather     string `json:"weather"`
}

// Openweather APP ID
const weatherAppID = "d4b24c62a0bc06b7238d9c0aa8fb1f6d"
const celciusValue = 273.15

// Function to get weather data using the Openweather external API
// Created By Balangoda Pamodi Bhagya - 500229522
func getWeather(city string) (*Weather, error) {
	// Openweather API Url to fetch the current weather data
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, weatherAppID)

	// Send an HTTP GET request to Opoenweather Url
	response, err := http.Get(url)

	// Check if an error occurred during the request
	if err != nil {
		// Return nil for the response and the error
		return nil, err
	}

	// Defer the closing of the response body
	defer response.Body.Close()

	// Read the response body using the io.ReadAll function
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var weatherData map[string]interface{}

	// Unmarshal the JSON-encoded response body into the variable named as weatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, err
	}

	// Extract relevant weather information
	// Convert the temperature value to a floating-point number
	temperatureValue := weatherData["main"].(map[string]interface{})["temp"].(float64)

	// Deduct celciusValue from the temperature value to convert it to Celsius
	temperatureCelsius := temperatureValue - celciusValue

	// Format the temperature in Celsius with one decimal place and append the unit "°C"
	temperatureString := fmt.Sprintf("%.1f°C", temperatureCelsius)

	// Return a pointer to a new Weather struct instance initialized with the extracted data
	return &Weather{
		City:        city,
		Temperature: temperatureString,
	}, nil
}
