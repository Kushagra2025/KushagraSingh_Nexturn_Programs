package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define a struct to hold the climate data for each city
type CityData struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

// Function to find the city with the highest temperature
func highestTemperature(cities []CityData) CityData {
	var highest CityData
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// Function to find the city with the lowest temperature
func lowestTemperature(cities []CityData) CityData {
	var lowest CityData
	for _, city := range cities {
		if lowest.Temperature == 0 || city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// Function to calculate the average rainfall across all cities
func averageRainfall(cities []CityData) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter cities by rainfall threshold
func filterCitiesByRainfall(cities []CityData, threshold float64) []CityData {
	var filteredCities []CityData
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

// Function to search for a city by name
func searchCityByName(cities []CityData, cityName string) (CityData, bool) {
	for _, city := range cities {
		if strings.ToLower(city.Name) == strings.ToLower(cityName) {
			return city, true
		}
	}
	return CityData{}, false
}

func main_5() {
	// Hardcoded data for demonstration purposes
	cities := []CityData{
		{"New York", 12.5, 120.0},
		{"London", 15.0, 800.0},
		{"Tokyo", 18.0, 1000.0},
		{"Sydney", 22.5, 120.0},
		{"Rio de Janeiro", 27.0, 1450.0},
	}

	// 1. Find city with highest and lowest temperature
	highest := highestTemperature(cities)
	lowest := lowestTemperature(cities)

	// Output highest and lowest temperature cities
	fmt.Printf("The city with the highest temperature is: %s with %.2f°C\n", highest.Name, highest.Temperature)
	fmt.Printf("The city with the lowest temperature is: %s with %.2f°C\n", lowest.Name, lowest.Temperature)

	// 2. Calculate and display average rainfall
	avgRainfall := averageRainfall(cities)
	fmt.Printf("The average rainfall across all cities is: %.2f mm\n", avgRainfall)

	// 3. Filter cities by rainfall threshold
	var thresholdInput string
	fmt.Print("Enter rainfall threshold (in mm) to filter cities: ")
	_, err := fmt.Scan(&thresholdInput)
	if err != nil {
		fmt.Println("Error: Invalid input for threshold.")
		return
	}

	// Convert threshold to float
	threshold, err := strconv.ParseFloat(thresholdInput, 64)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid number for the threshold.")
		return
	}

	filteredCities := filterCitiesByRainfall(cities, threshold)
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	for _, city := range filteredCities {
		fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
	}

	// 4. Search for a city by name
	var cityName string
	fmt.Print("\nEnter a city name to search for: ")
	fmt.Scan(&cityName)

	city, found := searchCityByName(cities, cityName)
	if found {
		fmt.Printf("City found: %s\n", city.Name)
		fmt.Printf("Temperature: %.2f°C\n", city.Temperature)
		fmt.Printf("Rainfall: %.2f mm\n", city.Rainfall)
	} else {
		fmt.Println("City not found.")
	}
}
