package main

import "fmt"

func main() {
	var choice int
	var temp float64

	fmt.Println("1 - Celsius to Fahrenheit")
	fmt.Println("2 - Fahrenheit to Celsius")
	fmt.Print("Choose option: ")
	fmt.Scanln(&choice)

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temp)

	switch choice {
	case 1:
		result := cToF(temp)
		fmt.Printf("Fahrenheit: %.2f\n", result)
	case 2:
		result := fToC(temp)
		fmt.Printf("Celsius: %.2f\n", result)
	default:
		fmt.Println("Invalid option")
	}
}

func cToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}