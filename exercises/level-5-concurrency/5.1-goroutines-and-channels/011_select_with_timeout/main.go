package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select with Timeout ===")

	ch := make(chan int)

	// Simula envio lento
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
	}()

	for {
		select {

		case v, ok := <-ch:
			if !ok {
				fmt.Println("\nChannel closed.")
				return
			}
			fmt.Println("Received:", v)

		case <-time.After(2 * time.Second):
			fmt.Println("\nTimeout reached. Stopping consumer.")
			return
		}
	}
}
