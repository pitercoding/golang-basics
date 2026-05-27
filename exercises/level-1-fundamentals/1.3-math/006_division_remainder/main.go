package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return
	}

	remainder := a % b

	fmt.Printf("Remainder: %d\n", remainder)
}