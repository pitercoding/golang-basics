package main

import "fmt"

func main() {
	var text1 string
	var text2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&text1)

	fmt.Print("Enter second string: ")
	fmt.Scanln(&text2)

	compareStrings(text1, text2)
}

func compareStrings(text1 string, text2 string) {
	if text1 == text2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are different")
	}
}