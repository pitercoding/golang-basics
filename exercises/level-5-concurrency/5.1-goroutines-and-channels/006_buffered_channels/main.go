package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n=== Buffered Channel ===")

	numbers := make(chan int, 3)

	go producer(numbers)
	go consumer(numbers)

	time.Sleep(3 * time.Second)

	fmt.Println("Done!")
}

func producer(numbers chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sending: %d\n", i)
		numbers <- i
		time.Sleep(300 * time.Millisecond)
	}

	close(numbers)
}

func consumer(numbers chan int) {
	for number := range numbers {
		fmt.Printf("Received: %d\n", number)
		time.Sleep(600 * time.Millisecond)
	}
}
