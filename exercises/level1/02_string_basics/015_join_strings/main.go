package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Go", "is", "awesome"}
	names := []string{"Ana", "Pia", "Bia"}

	resultWords := strings.Join(words, " ")
	resultNames := strings.Join(names, ", ")

	fmt.Println(resultWords)
	fmt.Println(resultNames)
}