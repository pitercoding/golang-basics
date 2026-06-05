package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var radius float64

	fmt.Println("\n=== Compute Circle Area ===")

	fmt.Print("Radius: ")
	fmt.Scanln(&radius)

	if radius <= 0 {
		fmt.Println("Radius must be greater than zero!")
		return
	}

	circle := Circle{
		Radius: radius,
	}

	fmt.Printf("\nCircle Area: %.2f\n", circle.Area())
	fmt.Printf("Circumference: %.2f\n", circle.Circumference())
}
