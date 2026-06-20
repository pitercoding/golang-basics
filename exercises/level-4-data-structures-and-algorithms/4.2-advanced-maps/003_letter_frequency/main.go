package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is awesome go go"

	result := letterFrequency(text)

	fmt.Println("Text:", text)
	fmt.Println("Frequency:", result)
}

func letterFrequency(text string) map[string]int {
	freq := make(map[string]int)

	text = strings.ReplaceAll(text, " ", "")

	for _, char := range text {
		letter := string(char)
		freq[letter]++
	}

	return freq
}