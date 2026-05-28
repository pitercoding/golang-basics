package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Table of %d:\n", i)

		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
