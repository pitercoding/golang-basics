package main

func main() {
	calculadora := &Calculadora{}
	runner := NewRunner(calculadora)
	runner.Execute()
}