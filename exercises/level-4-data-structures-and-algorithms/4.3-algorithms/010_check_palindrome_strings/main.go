package main

import "fmt"

func main() {
	words := []string{
		"radar",
		"level",
		"madam",
		"hello",
		"golang",
	}

	fmt.Println("\n=== Check Palindrome Strings ===")

	for _, word := range words {
		fmt.Printf("%s -> %t\n", word, isPalindrome(word))
	}

}

func isPalindrome(text string) bool {
	left := 0
	right := len(text) - 1

	for left < right {
		if text[left] != text[right] {
			return false
		}

		left++
		right--
	}

	return true
}
