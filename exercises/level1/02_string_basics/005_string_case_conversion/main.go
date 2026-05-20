package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Golang is Awesome!"

	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)

	fmt.Println("Original:", text)
	fmt.Println("Uppercase:", upper)
	fmt.Println("Lowercase:", lower)
}