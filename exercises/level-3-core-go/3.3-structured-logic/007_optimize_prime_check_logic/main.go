package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("\n=== Prime Number Checker ===")

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if isPrime(number) {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a prime number")
	}

	fmt.Printf("\nListing all prime numbers until %d:\n", number)
	printPrime(number)
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func printPrime(number int)  {
	for i := 2; i <= number; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}