package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID    int
	Value int
}

type Result struct {
	TaskID int
	Output int
	Worker int
}

func worker(id int, jobs <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing task %d\n", id, job.ID)

		time.Sleep(300 * time.Millisecond)

		results <- Result{
			TaskID: job.ID,
			Output: job.Value * 2,
			Worker: id,
		}
	}
}

func main() {
	fmt.Println("\n=== Worker Pool Task Runner ===")

	const workerCount = 3
	const taskCount = 10

	jobs := make(chan Task, taskCount)
	results := make(chan Result, taskCount)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send tasks
	for i := 1; i <= taskCount; i++ {
		jobs <- Task{
			ID:    i,
			Value: i,
		}
	}

	close(jobs)

	// Wait workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf(
			"Task %d processed by Worker %d -> Output: %d\n",
			result.TaskID,
			result.Worker,
			result.Output,
		)
	}

	fmt.Println("\nAll tasks completed.")
}
