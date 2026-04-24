package main

import "fmt"

func alterarCopia (x int) int {
	fmt.Println("Dentro da função. Definido como:", x)
	x = x * 2
	fmt.Println("Dentro da função. Atualizado para:", x)
	return x
}

func main() {
	numero := 10
	fmt.Println("Fora da função. Definido como:", numero)
	y := alterarCopia(numero)
	fmt.Println("Fora da função. Atualizado para:", y)
}