package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Enter radius: ")
	fmt.Scanln(&radius)

	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("Area of the circle: %.2f\n", area)
}