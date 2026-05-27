package main

import "fmt"

func main() {
	var base int
	var exponent int

	fmt.Print("Enter base: ")
	fmt.Scanln(&base)

	fmt.Print("Enter exponent: ")
	fmt.Scanln(&exponent)

	if exponent < 0 {
		fmt.Println("Negative exponents not supported")
		return
	}

	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}

	fmt.Printf("%d ^ %d = %d\n", base, exponent, result)
}
