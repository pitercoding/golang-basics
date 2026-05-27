package main

import "fmt"

func main() {
	var base float64
	var height float64

	fmt.Print("Enter base: ")
	fmt.Scanln(&base)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	area := calculateArea(base, height)

	fmt.Printf("Area of rectangle: %.2f\n", area)
}

func calculateArea(base, height float64) float64 {
	return base * height
}