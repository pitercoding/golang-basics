package main

import "fmt"

func main() {
	number := 12345
	count := 0

	for number != 0 {
		number /= 10
		count++
	}

	fmt.Printf("Number of digits: %d\n", count)
}
