package main

import "fmt"

func main() {
	var n1, n2, n3 float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&n1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&n2)

	fmt.Print("Enter third number: ")
	fmt.Scanln(&n3)

	average := (n1 + n2 + n3) / 3

	fmt.Printf("Average: %.2f\n", average)
}