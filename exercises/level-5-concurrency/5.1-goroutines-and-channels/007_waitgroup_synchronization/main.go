package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("\n=== WaitGroup Example ===")

	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()

	fmt.Println("\nAll goroutines finished.")
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()

	for letter := 'A'; letter <= 'E'; letter++ {
		fmt.Printf("%c\n", letter)
		time.Sleep(300 * time.Millisecond)
	}
}