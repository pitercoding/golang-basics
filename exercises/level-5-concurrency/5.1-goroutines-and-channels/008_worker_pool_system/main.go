package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Worker Pool ===")

	const workers = 3
	const totalJobs = 6

	jobs := make(chan int)

	var wg sync.WaitGroup

	for workerID := 1; workerID <= workers; workerID++ {
		wg.Add(1)
		go worker(workerID, jobs, &wg)
	}

	for job := 1; job <= totalJobs; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()

	fmt.Println("\nAll jobs completed.")
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
	}
}