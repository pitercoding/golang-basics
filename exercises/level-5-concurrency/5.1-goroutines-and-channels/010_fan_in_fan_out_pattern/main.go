package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Fan-in / Fan-out ===")

	jobs := generator(1, 10)

	// FAN-OUT (workers)
	w1 := worker(1, jobs)
	w2 := worker(2, jobs)
	w3 := worker(3, jobs)

	// FAN-IN (merge results)
	results := fanIn(w1, w2, w3)

	for r := range results {
		fmt.Println("Result:", r)
	}

	fmt.Println("\nDone!")
}

func generator(start, end int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := start; i <= end; i++ {
			out <- i
		}
	}()

	return out
}

func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			result := v * v
			fmt.Printf("Worker %d processed %d -> %d\n", id, v, result)
			out <- result
		}
	}()

	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()

			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
