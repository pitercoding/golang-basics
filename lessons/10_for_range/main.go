package main
import "fmt"

func main() {
	numeros := []int{10, 20, 30}
	for i, numero := range numeros {
		fmt.Printf("Elemento %d: %d\n", i, numero)
	}
}