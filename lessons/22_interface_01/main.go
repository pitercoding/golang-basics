package main

import "fmt"

type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s está acelerando!\n", c.Modelo)
}

type Bicileta struct {
	Tipo string
}

func (b Bicileta) Acelerar() {
	fmt.Printf("A bicicleta do tipo %s está acelerando!\n", b.Tipo)
}

// Definindo uma interface para comportamentos
type Veiculo interface {
	Acelerar()
}

func TestarVeiculo(v Veiculo) {
	v.Acelerar()
}

func main() {
	carro := Carro{Modelo: "Sedan"}
	bike := Bicileta{Tipo: "Mountain Bike"}
	TestarVeiculo(carro)
	TestarVeiculo(bike)
}