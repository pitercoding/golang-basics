package main

import (
	"fmt"
)

func main() {
	var option int

	for {
		fmt.Println("\n=== Unit Converter ===")
		fmt.Println("1. Kilometers to Meters")
		fmt.Println("2. Meters to Kilometers")
		fmt.Println("3. Celsius to Fahrenheit")
		fmt.Println("4. Fahrenheit to Celsius")
		fmt.Println("5. Minutes to Hours")
		fmt.Println("6. Hours to Minutes")
		fmt.Println("0. Exit")

		fmt.Print("Choose your option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var km float64

			fmt.Print("Enter kilometers: ")
			fmt.Scan(&km)

			result := kilometersToMeters(km)

			fmt.Printf("%.2f km = %.2f m\n", km, result)

		case 2:
			var meters float64

			fmt.Print("Enter meters: ")
			fmt.Scan(&meters)

			result := metersToKilometers(meters)

			fmt.Printf("%.2f m = %.2f km\n", meters, result)

		case 3:
			var celsius float64

			fmt.Print("Enter °C: ")
			fmt.Scan(&celsius)

			result := celsiusToFahrenheit(celsius)

			fmt.Printf("%.2f°C = %.2f°F\n", celsius, result)

		case 4:
			var fahrenheit float64

			fmt.Print("Enter °F: ")
			fmt.Scan(&fahrenheit)

			result := fahrenheitToCelsius(fahrenheit)

			fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, result)

		case 5:
			var minutes float64

			fmt.Print("Enter minutes: ")
			fmt.Scan(&minutes)

			result := minutesToHours(minutes)

			fmt.Printf("%.2f min = %.2f h\n", minutes, result)

		case 6:
			var hours float64

			fmt.Print("Enter hours: ")
			fmt.Scan(&hours)

			result := hoursToMinutes(hours)

			fmt.Printf("%.2f h = %.2f min\n", hours, result)

		case 0:
			fmt.Println("\nThank you for your time. Exiting program.")
			return
		default:
			fmt.Println("\nInvalid option! Try again.")

		}
	}
}

func kilometersToMeters(km float64) float64 {
	return km * 1000
}

func metersToKilometers(meters float64) float64 {
	return meters / 1000
}

func celsiusToFahrenheit(celcius float64) float64 {
	return (celcius * 1.8) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

func minutesToHours(minutes float64) float64 {
	return minutes / 60
}

func hoursToMinutes(hours float64) float64 {
	return hours * 60
}
