package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main started.")

	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)

	fmt.Println("Main finished.")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters() {
	for letter := 'A'; letter <= 'E'; letter++ {
		fmt.Printf("%c\n", letter)
		time.Sleep(300 * time.Millisecond)
	}
}
