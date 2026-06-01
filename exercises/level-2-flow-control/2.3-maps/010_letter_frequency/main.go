package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "GoLAng"

	clean := strings.ToLower(text)

	freq := map[rune]int{}

	for _, char := range clean {
		freq[char]++
	}

	for char, count := range freq {
		fmt.Printf("%c -> %d\n", char, count)
	}
}
