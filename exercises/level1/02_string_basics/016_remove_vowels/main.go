package main

import "fmt"

func main() {
	text := "Golang is awesome"

	result := ""

	for _, letter := range text {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			continue
		}

		result += string(letter)
	}

	fmt.Println("Original:", text)
	fmt.Println("Without vowels:", result)

}