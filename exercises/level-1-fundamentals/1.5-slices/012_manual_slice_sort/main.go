package main

import "fmt"

func main() {
	numbers := []int{5, 2, 8, 1, 9}

	bubbleSort(numbers)

	fmt.Println("Sorted slice:", numbers)
}

func bubbleSort(slice []int) {
	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
