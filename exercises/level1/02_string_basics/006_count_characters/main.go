package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	word1 := "golang"
	word2 := "olá!"

	text := "Go Lang Basics"
	noSpaces := strings.ReplaceAll(text, " ", "")


	length1 := len(word1)
	length2 := utf8.RuneCountInString(word2)
	length3 := utf8.RuneCountInString(noSpaces)

	fmt.Println("Total characters:", length1)
	fmt.Println("Total characters:", length2)
	fmt.Println("Total characters:", length3)
}