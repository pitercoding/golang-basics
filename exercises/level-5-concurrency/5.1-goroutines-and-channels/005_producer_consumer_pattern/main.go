package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Producer Consumer ===")

	numbers := make(chan int)

	go producer(numbers)
	go consumer(numbers)

	time.Sleep(2 * time.Second)

	fmt.Println("\nDone!")
}

func producer(numbers chan int)  {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Produced: %d\n", i)
		numbers <- i
		time.Sleep(300 * time.Millisecond)
	}

	close(numbers)
}

func consumer(numbers chan int)  {
	for number := range numbers {
		fmt.Printf("Consumed: %d\n", number)
	}
}