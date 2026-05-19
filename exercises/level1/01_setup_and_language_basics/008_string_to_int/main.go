package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid number!")
		return
	}

	
	fmt.Printf("Converted number: %d\n", number)
}