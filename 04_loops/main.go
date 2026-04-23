package main

import "fmt"

func Main1() {
	for i := 1; i < 5; i++ {
		fmt.Printf("Iteration number: %d\n", i)
	}
}

func main() {
	var age int
	for {
		fmt.Print("Enter your age (must be at least 18): ")
		_, err := fmt.Scanln(&age)

		if err != nil {
			fmt.Println("Invalid input! Please try again.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if age < 18 {
			fmt.Println("Invalid input! Please try again.")
			continue
		}
		break
	}
	fmt.Println("Access granted. Welcome!")
}