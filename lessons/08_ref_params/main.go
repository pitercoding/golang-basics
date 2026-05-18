package main
import "fmt"

func alterarOriginal (x *int) {
	*x = *x * 2
}

func main() {
	numero := 10
	fmt.Println("Fora da função. Definido como:", numero)
	alterarOriginal(&numero)
	fmt.Println("Fora da função. Atualizado para:", numero)
}