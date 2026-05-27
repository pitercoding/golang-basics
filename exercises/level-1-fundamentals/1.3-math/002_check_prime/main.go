package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	if n <= 1 {
		fmt.Println("Not a prime number")
		return
	}

	isPrime := true

	for i := 2; i * i <= n; i++ {
		if  n % i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d is a prime number", n)
	} else {
		fmt.Printf("%d is not a prime number", n)
	}
}