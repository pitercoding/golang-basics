package main

import "fmt"

func main() {
	numbers := []int{5, 10, 15, 20, 25}

	for _, n := range numbers {
		fmt.Println(n)
	}
}