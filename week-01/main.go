package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertCelciusToReamur(celcius float32) float32 {
	return celcius * 0.8
}

func convertCelciusToFahrenheit(celcius float32) float32 {
	return (celcius * 1.8) + 32
}

func main() {
	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan suhu dalam Celsius: ")

	reader := bufio.NewReader(os.Stdin)
	inputValue, _ := reader.ReadString('\n')

	inputValue = strings.TrimSpace(inputValue)
	convertToFloat, err := strconv.ParseFloat(inputValue, 32)

	// Validasi jika input tidak sesuai
	if err != nil {
		fmt.Println("Input Tidak Valid, hanya menerima angka")
		return
	}

	celciusValue := float32(convertToFloat)

	// Convert & Print Reamur
	reamurValue := convertCelciusToReamur(celciusValue)
	fmt.Println("Suhu dalam Reamur =", reamurValue)

	// Convert & Print Fahrenheit
	fahrenheitValue := convertCelciusToFahrenheit(celciusValue)
	fmt.Println("Suhu dalam Fahrenheit =", fahrenheitValue)
}
