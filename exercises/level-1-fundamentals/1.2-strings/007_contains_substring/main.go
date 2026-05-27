package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "golang"
	substring := "go"

	contains := strings.Contains(text, substring)

	fmt.Println("Contains substring:", contains)
}