package main

import (
	"fmt"
)

func main() {
	var number float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Square root not defined for negative numbers")
		return
	}

	guess := number / 2

	for i := 0; i < 10; i++ {
		guess = (guess + number/guess) / 2
	}

	fmt.Printf("Approximate square root: %.5f\n", guess)
}
