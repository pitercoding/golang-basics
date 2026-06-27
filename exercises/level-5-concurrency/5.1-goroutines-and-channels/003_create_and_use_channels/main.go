package main

import "fmt"

func main() {
	fmt.Println("Main started.")

	messages := make(chan string)

	go sendMessage(messages)

	message := <- messages

	fmt.Println("Received:", message)
	fmt.Println("Main finished.")
}

func sendMessage(messages chan string) {
	fmt.Println("Sending message...")

	messages <- "Hello from the goroutine!"
}