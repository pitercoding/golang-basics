package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "golang"

	result := strings.ReplaceAll(text, "g", "G")

	fmt.Println("Original:", text)
	fmt.Println("Modified:", result)
}