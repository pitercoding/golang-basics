package main

import "fmt"

func main() {
	var n1 int
	var n2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&n1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&n2)

	sum := n1 + n2

	fmt.Printf("Sum: %d\n", sum)
}