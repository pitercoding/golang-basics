package main

import "fmt"

// Comportamento com interface
type Acelerador interface{
	Acelerar()
}

// Comportamento com interface
type Freio interface{
	Frear()
}

// Struct base: Combinando comportamentos
type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s está acelerando!\n", c.Modelo)
}

func (c Carro) Frear() {
	fmt.Printf("O carro %s está freando!\n", c.Modelo)
}

// Adicionando novos comportamentos por composição
type CarroEletrico struct {
	Carro //Composição
	BateriaCarga int
}

func (ce CarroEletrico) CarregarBateria() {
	fmt.Printf("Carregando bateria... Carga: %d%%\n", ce.BateriaCarga)
}

func main() {
	ce := CarroEletrico {
		Carro: Carro{Modelo: "Tesla"},
		BateriaCarga: 80,
	}

	ce.Acelerar() // Método herdado por composição
	ce.Frear() // Método herdado por composição
	ce.CarregarBateria() // Método especifico de CE
}