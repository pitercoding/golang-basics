package main

import "fmt"

func sum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var num1, num2 int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter another number: ")
	fmt.Scanln(&num2)

	sum := sum(num1, num2)

	fmt.Printf("The sum of %d and %d is %d \n", num1, num2, sum)
}