package main

import "fmt"

type Pessoa struct {
	Nome  string
	Maior bool
	Idade int
}

func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, eu sou %s e tenho %d anos.", p.Nome, p.Idade)
}

func (p *Pessoa) Envelhecer() {
	p.Idade++
}

func main() {
	pessoa := Pessoa{"Racha Cuca", true, 38}
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Maior:", pessoa.Maior)
	fmt.Println("Idade:", pessoa.Idade)
	fmt.Println("pessoa:", pessoa)
	fmt.Println("Pessoa &:", &pessoa)
	fmt.Println("Pessoa.Nome &:", &pessoa.Nome)
	fmt.Println("Pessoa.Maior &:", &pessoa.Maior)
	fmt.Println("Pessoa.Idade &:", &pessoa.Idade)
	fmt.Println("Racha Cuca diz:", pessoa.Apresentar())

	pessoa.Envelhecer()
	fmt.Println("Idade:", pessoa.Idade)
}