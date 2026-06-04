package main

import "fmt"

func main() {
	var n int

	fmt.Println("\n=== Recursive Fibonacci ===")

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	if n < 0 {
		fmt.Println("Fibonacci is not defined for negative numbers")
		return
	}

	result := fibonacci(n)

	fmt.Printf("F(%d) = %d\n", n, result)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
