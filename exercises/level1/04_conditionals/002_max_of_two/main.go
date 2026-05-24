package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	max := findMax(a, b)

	fmt.Printf("Maximum number: %.2f\n", max)
}

func findMax(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}