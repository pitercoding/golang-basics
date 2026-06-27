package main

import "fmt"

func main() {
	fmt.Println("\n=== Channel Example ===")

	numbers := make(chan int)

	go produceNumbers(numbers)

	for i := 0; i < 5; i++ {
		number := <- numbers
		fmt.Printf("Received: %d\n", number)
	}

	fmt.Println("\nDone!")
}

func produceNumbers(numbers chan int)()  {
	for i := 1; i <= 5; i++ {
		numbers <- i
	} 
}