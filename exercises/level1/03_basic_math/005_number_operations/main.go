package main

import "fmt"

func main() {
	var n float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	double := n * 2
	triple := n * 3
	half := n / 2

	fmt.Printf("Number: %.2f\n", n)
	fmt.Printf("Double: %.2f\n", double)
	fmt.Printf("Triple: %.2f\n", triple)
	fmt.Printf("Half: %.2f\n", half)
}