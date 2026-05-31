package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 28,
		"David": 45,
		"Pablo": 21,
		"Ana": 17,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	fmt.Println(ages["Pablo"])
}
