package main

import "fmt"

func main() {
	var km float64

	fmt.Print("Enter kilometers: ")
	fmt.Scanln(&km)

	fmt.Printf("%.2f km = %.2f meters\n", km, kmToMeters(km))
}

func kmToMeters(km float64) float64 {
	const metersPerKm = 1000
	return km * metersPerKm
}