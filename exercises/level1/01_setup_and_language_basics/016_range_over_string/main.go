package main

import "fmt"

func main() {
	word := "Golang"

	for index, letter := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, letter)
	}
}