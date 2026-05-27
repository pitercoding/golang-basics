package main

import "fmt"

func main() {
	var age int
	var price float64

	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter ticket price: ")
	fmt.Scanln(&price)

	finalPrice := applyDiscount(age, price)

	fmt.Printf("Final price: %.2f\n", finalPrice)
}

func applyDiscount(age int, price float64) float64 {
	
	if age < 12 {
		return price * 0.5
	}

	if age >= 60 {
		return price  * 0.7
	}
	
	return price
}