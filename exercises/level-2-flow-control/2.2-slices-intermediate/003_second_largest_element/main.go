package main

import "fmt"

func main() {
	numbers := []int{10, 5, 20, 8, 15}

	largest := numbers[0]
	secondLargest := numbers[1]

	if secondLargest > largest {
		largest, secondLargest = secondLargest, largest
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > largest {
			secondLargest = largest
			largest = numbers[i]
		} else if numbers[i] > secondLargest {
			secondLargest = numbers[i]
		}
	}

	fmt.Printf("Second largest number: %d\n", secondLargest)
}