package main

import "fmt"

func main() {
	word := "golang"
	vowelCount := 0

	for _, letter := range word {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelCount++
		}
	}

	fmt.Printf("Vowels: %d\n", vowelCount)
}