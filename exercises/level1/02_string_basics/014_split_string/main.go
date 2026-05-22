package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "apple,banana,orange"

	fruits := strings.Split(text, ",")

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
}