package main

import "fmt"

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	var width float64
	var height float64

	fmt.Println("\n=== Rectangle Area Calculator ===")

	fmt.Print("Enter width: ")
	fmt.Scanln(&width)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	if width <= 0 || height <= 0 {
		fmt.Println("Width and height must be greater than 0")
		return
	}

	rectangle := Rectangle {
		Width: width,
		Height: height,
	}

	fmt.Printf("\nRectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle.Perimeter())
}