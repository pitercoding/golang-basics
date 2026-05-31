package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunkSize := 3

	var chunks [][]int

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize

		if end > len(numbers) {
			end = len(numbers)
		}

		chunks = append(chunks, numbers[i:end])
	}

	fmt.Println("Chunks:")

	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
}
