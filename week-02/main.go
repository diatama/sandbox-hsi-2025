package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WeatherStruct struct {
	Location string
	Temperature float32
	Condition string
}

func convertCelciusToReamur(Weather *WeatherStruct) float32 {
	return Weather.Temperature * 0.8
}

func convertCelciusToFahrenheit(Weather *WeatherStruct) float32 {
	return (Weather.Temperature * 1.8) + 32
}

func ConvertCelciusToCondition(Weather *WeatherStruct) string {
	if Weather.Temperature < 18 {
		return "Dingin"
	} else if Weather.Temperature >= 18 && Weather.Temperature <= 25 {
		return "Hangat"
	}
	return "Panas"
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	var Weather WeatherStruct

	// Open title
	fmt.Println("--- Konverter Suhu ---")
	
	// Weather Location
	fmt.Print("Masukkan lokasi pengukuran suhu: ")
	inputLocation, _ := reader.ReadString('\n')
	inputLocation = strings.TrimSpace(inputLocation)
	_, err := strconv.Atoi(inputLocation)

	// Validasi Location adalah string
	if( err == nil ) {
		fmt.Println("Input lokasi hanya menerima string.")
		return
	}

	Weather.Location = inputLocation
	
	// Weather Temperature
	fmt.Print("Masukkan suhu dalam Celsius: ")
	inputTemperature, _ := reader.ReadString('\n')
	inputTemperature = strings.TrimSpace(inputTemperature)
	convertToFloat, err := strconv.ParseFloat(inputTemperature, 32)

	// Validasi Temperature adalah angka
	if err != nil {
		fmt.Println("Input suhu, hanya menerima angka")
		return
	}

	Weather.Temperature = float32(convertToFloat)


	fmt.Println("Suhu di", Weather.Location, "adalah", ConvertCelciusToCondition(&Weather))

	// Convert & Print Reamur
	reamurValue := convertCelciusToReamur(&Weather)
	fmt.Println("Suhu di", Weather.Location, "dalam Reamur =", reamurValue)

	// Convert & Print Fahrenheit
	fahrenheitValue := convertCelciusToFahrenheit(&Weather)
	fmt.Println("Suhu di", Weather.Location, "dalam Fahrenheit =", fahrenheitValue)
}
