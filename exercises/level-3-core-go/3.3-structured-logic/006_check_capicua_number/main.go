package main

import "fmt"

func main() {
	var number int

	fmt.Println("\n=== Capicua Number Checker ===")

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers are not supported.")
		return
	}

	if isCapicua(number) {
		fmt.Println("The number is capicua.")
	} else {
		fmt.Println("The number is not capicua.")
	}
}

func isCapicua(number int) bool {
	return number == reverseNumber(number)
}

func reverseNumber(number int) int {
	reversed := 0

	for number > 0 {
		digit := number % 10

		reversed = reversed*10 + digit

		number /= 10
	}

	return reversed
}