package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Print("Enter first side: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second side: ")
	fmt.Scanln(&b)

	hypotenuse := calculateHypotenuse(a, b)

	fmt.Printf("Hypotenuse: %.2f\n", hypotenuse)
}

func calculateHypotenuse(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}
