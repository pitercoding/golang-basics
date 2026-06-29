package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Context Cancellation ===")

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)
	go worker(ctx, 3)

	time.Sleep(2 * time.Second)

	fmt.Println("\n>>> Cancelling execution...")
	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Done!")
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return

		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
