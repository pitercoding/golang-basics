package main

import (
	"errors"
	"fmt"
)

type Runner struct {
	calculadora *Calculadora
	operacao string
}

func (r *Runner) SolicitarValores() error {
	fmt.Print("Primeiro valor:")
	if _, err := fmt.Scanln(&r.calculadora.Operando1); err != nil {
		return errors.New("Entrada inválida para primeiro valor!")
	}

	fmt.Print("Segundo valor:")
	if _, err := fmt.Scanln(&r.calculadora.Operando2); err != nil {
		return errors.New("Entrada inválida para segundo valor!")
	}

	return nil
}

func (r *Runner) SolicitarOperacao() error {
	var operacao string

	fmt.Print("Escolha a operação (+, -, *, /):")
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("Entrada inválida para operação!")
	}

	switch operacao {
	case "+", "-", "*", "/":
		r.operacao = operacao
		return nil
	default:
		return errors.New("Operação inválida!")
	}
}

func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
	case "+":
		fmt.Println("Resultado:", r.calculadora.Somar())
	case "-":
		fmt.Println("Resultado:", r.calculadora.Subtrair())
	case "*":
		fmt.Println("Resultado:", r.calculadora.Multiplicar())
	case "/":
		resultado, err := r.calculadora.Dividir()
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Resultado:", resultado)
		}
	}
}

func (r *Runner) Execute() {
	for {
		if err := r.SolicitarValores(); err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		err := r.SolicitarOperacao()
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		r.ExecutarOperacao()
	}
}

func NewRunner(c *Calculadora) *Runner {
	return &Runner{calculadora: c}
}