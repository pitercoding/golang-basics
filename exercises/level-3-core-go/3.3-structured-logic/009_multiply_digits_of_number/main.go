package main

import "fmt"

func main() {
	var number int

	fmt.Println("\n=== Multiply Digits ===")

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers are not supported.")
		return
	}

	result, count := multiplyAndCountDigits(number)

	fmt.Printf("Product of digits: %d\n", result)
	fmt.Printf("Number of digits: %d\n", count)
}

func multiplyAndCountDigits(number int) (int, int) {
	if number == 0 {
		return 0, 1
	}

	product := 1
	count := 0

	for number > 0 {
		digit := number % 10

		product *= digit

		number /= 10
		count++
	}

	return product, count
}
