package main

import "fmt"

func main() {
	var n int

	fmt.Println("\n=== Recursive Factorial ===")

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers")
		return
	}

	result := factorial(n)

	fmt.Printf("%d! = %d\n", n, result)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
