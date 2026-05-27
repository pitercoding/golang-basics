package main

import "fmt"

func main() {
	word := "radar"
	reversed := ""

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	if word == reversed {
		fmt.Println("Palindrome!")
	} else {
		fmt.Println("Not a palindrome!")
	}
}