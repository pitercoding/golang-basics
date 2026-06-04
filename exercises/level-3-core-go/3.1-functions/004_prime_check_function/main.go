package main

import "fmt"

func main() {
	var n int

	fmt.Println("\n=== Prime Check Function ===")

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Printf("%d is a prime number\n", n)
	} else {
		fmt.Printf("%d is NOT a prime number\n", n)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
