package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) Display() {
	fmt.Printf("Brand: %s\n", c.Brand)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
}

func main() {
	var brand, model string
	var year int

	fmt.Println("\n=== Car Struct ===")

	fmt.Print("Enter brand: ")
	fmt.Scanln(&brand)

	fmt.Print("Enter model: ")
	fmt.Scanln(&model)

	fmt.Print("Enter year: ")
	fmt.Scanln(&year)

	if year < 1866 {
		fmt.Println("Invalid year")
		return
	}

	car := Car {
		Brand: brand,
		Model: model,
		Year: year,
	}

	fmt.Println("\n--- Car Information ---")
	car.Display()
}