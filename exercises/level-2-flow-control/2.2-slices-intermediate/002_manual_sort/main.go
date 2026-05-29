package main

import "fmt"

func main() {
	arr := []int{5, 2, 9, 1, 5}

	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted slice:", arr)
}
