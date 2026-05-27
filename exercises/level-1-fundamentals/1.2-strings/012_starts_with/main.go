package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Golang"

	letter := "A"

	result := strings.HasPrefix(text, letter)

	fmt.Println("Starts with letter:", result)
}