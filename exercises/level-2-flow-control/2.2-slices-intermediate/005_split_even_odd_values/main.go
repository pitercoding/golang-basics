package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	even := []int{}
	odd := []int{}

	for _, n := range numbers {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	fmt.Println("Original Slice: ", numbers)
	fmt.Println("Even Numbers: ", even)
	fmt.Println("Odd Numbers: ", odd)
}
