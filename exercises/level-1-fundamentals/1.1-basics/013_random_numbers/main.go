package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)

	fmt.Println("Random Number:", randomNumber)
}