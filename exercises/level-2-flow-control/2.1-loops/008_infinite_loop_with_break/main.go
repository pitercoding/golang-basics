package main

import "fmt"

func main() {
	var number int

	for {
		fmt.Print("Enter a number (0 to stop): ")
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		fmt.Println("You entered:", number)
	}

	fmt.Println("Loop finished!")
}
