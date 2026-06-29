package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Synchronize Multiple Goroutines ===")

	var wg sync.WaitGroup

	totalWorkers := 10

	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("\nAll workers finished.")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)

	time.Sleep(time.Duration(id) * 500 * time.Millisecond)

	fmt.Printf("Worker %d finished\n", id)
}
