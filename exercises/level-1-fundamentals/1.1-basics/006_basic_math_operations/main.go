package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num1 / num2

	fmt.Printf("Subtraction: %.2f\n", subtraction)
	fmt.Printf("Multiplication: %.2f\n", multiplication)
	fmt.Printf("Division: %.2f\n", division)
}