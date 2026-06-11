package main

import "fmt"

func main() {
	var number int

	fmt.Println("\n=== Sum Digits ===")

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers are not supported.")
		return
	}

	result, count := sumAndCountDigits(number)

	fmt.Printf("Sum of digits: %d\n", result)
	fmt.Printf("Number of digits: %d\n", count)
}

func sumAndCountDigits(number int) (int, int) {
	if number == 0 {
		return 0, 1
	}

	sum := 0
	count := 0

	for number > 0 {
		digit := number % 10

		sum += digit

		number /= 10
		count++
	}

	return sum, count
}