package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Mutex Protection ===")

	var wg sync.WaitGroup
	var mutex sync.Mutex

	counter := 0

	totalWorkers := 100
	incrementsPerWorker := 1000

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < incrementsPerWorker; j++ {

				mutex.Lock()
				counter++
				mutex.Unlock()

			}
		}()
	}

	wg.Wait()

	expected := totalWorkers * incrementsPerWorker

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Actual:   %d\n", counter)
}