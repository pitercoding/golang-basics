package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	text := strconv.Itoa(number)

	fmt.Println("Converted string:", text)
}