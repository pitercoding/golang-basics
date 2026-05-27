package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	index := 2

	numbers = removeByIndex(numbers, index)

	fmt.Println("Updated slice:", numbers)
}

func removeByIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}