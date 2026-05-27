package main

import "fmt"

func main() {
	var signal int

	fmt.Print("Enter signal (1=Red, 2=Yellow, 3=Green): ")
	fmt.Scanln(&signal)
	
	printSignal(signal)
}

func printSignal(signal int) {
	switch signal {
	case 1:
		fmt.Println("Red - Stop")
	case 2:
		fmt.Println("Yellow - Attention")
	case 3:
		fmt.Println("Green - Go")
	default:
		fmt.Println("Invalid signal!")
	}
}