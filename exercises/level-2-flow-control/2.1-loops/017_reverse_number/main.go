package main

import "fmt"

func main() {
	n := 1234
	reversed := 0

	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n = n / 10
	}

	fmt.Printf("Reversed number: %d\n", reversed)
}
