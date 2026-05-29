package main

import "fmt"

func main() {
	n := 5
	result := 1

	for i := n; i >= 1; i-- {
		result *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, result)
}
