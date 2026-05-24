package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter third number: ")
	fmt.Scanln(&c)

	max := findMax(a, b, c)

	fmt.Printf("Maximum number: %.2f\n", max)
}

func findMax(a, b, c float64) float64 {
	max := a

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	return max
}
