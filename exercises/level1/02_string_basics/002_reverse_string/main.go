package main

import "fmt"

func main() {
	word := "golang"
	reversed := ""

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	fmt.Println("Original:", word)
	fmt.Println("Reversed:", reversed)
}