package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Pipeline Example ===")

	numbers := generator(1, 5)

	sqIn, cbIn := fanOut(numbers)

	squared := processorSquared(sqIn)
	cubed := processorCubed(cbIn)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range squared {
			fmt.Println("Squared:", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range cubed {
			fmt.Println("Cubed:", v)
		}
	}()

	wg.Wait()

	fmt.Println("\nDone!")
}

func generator(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i <= end; i++ {
			fmt.Println("Input:", i)
			out <- i
		}
	}()
	return out
}

func fanOut(in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range in {
			v1 := v
			v2 := v

			out1 <- v1
			out2 <- v2
		}
	}()

	return out1, out2
}

func processorSquared(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func processorCubed(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v * v
		}
	}()
	return out
}
