package main

import "fmt"

func main() {
	fmt.Println("=== Safely Close Channels ===")

	numbers := producer()

	for value := range numbers {
		fmt.Println("Received:", value)
	}

	fmt.Println("\nChannel closed safely.")
}

func producer() <-chan int {
	out := make(chan int)

	go func ()  {
		defer close(out)

		for i := 1; i <= 5; i++ {
			out <- i
		}
	}()

	return out
}