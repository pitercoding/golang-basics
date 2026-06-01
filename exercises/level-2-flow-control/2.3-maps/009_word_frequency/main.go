package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go, go GO python!, JAVA, TyPeSCRipt!"

	clean := strings.ToLower(
		strings.ReplaceAll(text, ",", ""),
	)

	clean = strings.ReplaceAll(clean, "!", "")

	words := strings.Split(clean, " ")

	freq := map[string]int{}

	for _, word := range words {
		freq[word]++
	}

	for word, count := range freq {
		fmt.Printf("%s -> %d\n", word, count)
	}
}
