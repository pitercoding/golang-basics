package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int

	fmt.Println("\n=== Numeric Palindrome Checker ===")

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers are not supported.")
		return
	}

	if isPalindrome(number) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
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