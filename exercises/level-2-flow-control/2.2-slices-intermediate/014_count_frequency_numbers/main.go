package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 3, 3, 3, 4}

	frequency := make(map[int]int)

	for _, number := range numbers {
		frequency[number]++
	}

	fmt.Println("Frequency count:")

	for number, count := range frequency {
		fmt.Printf("%d -> %d times\n", number, count)
	}
}
