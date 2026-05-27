package main

import "fmt"

func main() {
	text1 := "Golang"
	text2 := "Olá!"

	substring1 := text1[0:3]

	runes := []rune(text2)
	substring2 := string(runes[0:3])

	fmt.Println(substring1)
	fmt.Println(substring2)
}