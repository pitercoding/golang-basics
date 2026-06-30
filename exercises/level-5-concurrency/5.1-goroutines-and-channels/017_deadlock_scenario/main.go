package main

import "fmt"

func main() {
	fmt.Println("=== Deadlock Example ===")

	ch := make(chan int)

	fmt.Println("Waiting for value...")

	value := <-ch

	fmt.Println(value)
}
