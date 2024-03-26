package main

// created by Tejaswi Cheripally - 500229934
// Define a struct named Weather to represent weather-related data
type Weather struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	FeelsLike   string `json:"feelslike"`
	Weather     string `json:"weather"`
}
