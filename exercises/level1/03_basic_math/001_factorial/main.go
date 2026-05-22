package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	
	if n < 0 {
		fmt.Println("Factorial not defined for negative numbers")
		return
	}
	
	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, factorial)
}