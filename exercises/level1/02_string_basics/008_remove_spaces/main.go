package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go Lang Basics"

	noSpaces := strings.ReplaceAll(text, " ", "")

	fmt.Println("Original:", text)
	fmt.Println("Without spaces:", noSpaces)

}