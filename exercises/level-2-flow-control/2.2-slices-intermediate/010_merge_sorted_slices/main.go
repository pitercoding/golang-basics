package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}

	i, j := 0, 0
	merged := []int{}

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for i < len(a) {
		merged = append(merged, a[i])
		i++
	}

	for j < len(b) {
		merged = append(merged, b[j])
		j++
	}

	fmt.Println("Merged sorted slices:", merged)
}
