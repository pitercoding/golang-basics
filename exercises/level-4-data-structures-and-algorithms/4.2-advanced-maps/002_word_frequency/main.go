package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is awesome go go"

	result := wordFrequency(text)

	fmt.Println("Text:", text)
	fmt.Println("Frequency:", result)
}

func wordFrequency(text string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		freq[word]++
	}

	return freq
}
