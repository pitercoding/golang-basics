package main

import "fmt"

func main() {
	var side float64

	fmt.Print("Enter side length: ")
	fmt.Scanln(&side)

	perimeter := calculateSquarePerimeter(side)

	fmt.Printf("Square Perimeter: %.2f\n", perimeter)
}

func calculateSquarePerimeter(side float64) float64 {
	return 4 * side
}