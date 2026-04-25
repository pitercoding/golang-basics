package saudacao
import "fmt"

func sayHi(nome string) string {
	return fmt.Sprintf("Bem-vindo(a), %s!\n", nome)
}

func BoasVindas(nome string) {
	fmt.Println(sayHi(nome))
}
