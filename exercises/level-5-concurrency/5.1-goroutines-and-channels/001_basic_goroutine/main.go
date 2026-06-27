package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main started.")

	go printMessage()

	time.Sleep(2 * time.Second)

	fmt.Println("Main finished.")
}

func printMessage() {
	fmt.Println("Hello from the goroutine!")
}