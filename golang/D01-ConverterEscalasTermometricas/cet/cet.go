package main

import "fmt"

func main() {
	const BOILING_POINT_OF_WATER_IN_K = 373
	const KELVIN_TO_CELSIUS_CONVERSION_FACTOR = 273
	fmt.Println("Ponto de ebulição da água de Kelvin para Celsius: ")
	fmt.Printf("%d K - %d = %d C", BOILING_POINT_OF_WATER_IN_K, KELVIN_TO_CELSIUS_CONVERSION_FACTOR, BOILING_POINT_OF_WATER_IN_K-KELVIN_TO_CELSIUS_CONVERSION_FACTOR)
}
