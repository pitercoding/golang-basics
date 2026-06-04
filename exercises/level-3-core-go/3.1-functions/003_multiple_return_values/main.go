package main

import "fmt"

func main() {
	var n1, n2 float64

	fmt.Println("\n=== Multiple Return Values ===")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&n1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&n2)

	sum, sub := calculate(n1, n2)

	fmt.Printf("%.2f + %.2f = %.2f\n", n1, n2, sum)
	fmt.Printf("%.2f - %.2f = %.2f\n", n1, n2, sub)
}

func calculate(a, b float64) (float64, float64) {
	sum := a + b
	sub := a - b

	return sum, sub
}
