package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Race Condition ===")

	var wg sync.WaitGroup

	counter := 0

	totalWorkers := 100
	incrementsPerWorker := 1000

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < incrementsPerWorker; j++ {
				counter++
			}
		}()
	}

	wg.Wait()

	expected := totalWorkers * incrementsPerWorker

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Actual: %d\n", counter)
}