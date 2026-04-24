package main
import "fmt"

func Main1() {
	idade := 30 // Variável comum
	ponteiroIdade := &idade // Ponteiro para variável "idade"
	fmt.Println("Valor de idade:", idade) // 30
	fmt.Println("Endereço de idade:", ponteiroIdade) // Endereço
	fmt.Println("Valor via ponteiro:", *ponteiroIdade) // Endereço
}

func main() {
	numero := 42
	ponteiro := &numero
	fmt.Println("Antes da alteração:", numero) // 42
	*ponteiro = 99 // Altera o valor armazenado no endereço
	fmt.Println("Depois da alteração:", numero) // 99
}