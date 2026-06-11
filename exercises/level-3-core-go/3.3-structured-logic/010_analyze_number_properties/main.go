package main

import (
	"fmt"
	"strconv"
)

type NumberAnalysis struct {
	Number       int
	Even         bool
	Odd          bool
	Prime        bool
	Palindrome   bool
	DigitCount   int
	DigitSum     int
	DigitProduct int
}

func analyzeNumber(number int) NumberAnalysis {
	return NumberAnalysis{
		Number:       number,
		Even:         isEven(number),
		Odd:          isOdd(number),
		Prime:        isPrime(number),
		Palindrome:   isPalindrome(number),
		DigitCount:   countDigits(number),
		DigitSum:     sumDigits(number),
		DigitProduct: multiplyDigits(number),
	}
}

func main() {
	var number int

	fmt.Println("=== Number Analyzer ===")

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers are not supported.")
		return
	}

	result := analyzeNumber(number)
	fmt.Println("\n--- Number Analysis ---")
	fmt.Printf("Number: %d\n", result.Number)
	fmt.Printf("Even: %t\n", result.Even)
	fmt.Printf("Odd: %t\n", result.Odd)
	fmt.Printf("Prime: %t\n", result.Prime)
	fmt.Printf("Palindrome: %t\n", result.Palindrome)
	fmt.Printf("Digits: %d\n", result.DigitCount)
	fmt.Printf("Sum of digits: %d\n", result.DigitSum)
	fmt.Printf("Product of digits: %d\n", result.DigitProduct)
}

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
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

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)

	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func countDigits(number int) int {
	if number == 0 {
		return 1
	}

	count := 0

	for number > 0 {
		number /= 10
		count++
	}

	return count
}

func sumDigits(number int) int {
	sum := 0

	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}

func multiplyDigits(number int) int {
	if number == 0 {
		return 0
	}

	product := 1

	for number > 0 {
		product *= number % 10
		number /= 10
	}

	return product
}
