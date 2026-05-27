package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "email1234@test.com"

	count := 0

	for _, char := range text {
		if unicode.IsDigit(char) {
			count++
		}
	}

	fmt.Println("Total digits:", count)
}