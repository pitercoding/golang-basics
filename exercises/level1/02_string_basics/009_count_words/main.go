package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "I am learning Golang"

	words := strings.Fields(sentence)

	fmt.Println("Total words:", len(words))
}