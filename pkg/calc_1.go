package pkg

import "fmt"

type CalculadoraA struct {
}

var _ CalculadoraInterface = (*CalculadoraA)(nil)

func (c *CalculadoraA) Soma(a int, b int) int {
	fmt.Println("CalculadoraA.Soma")
	return a + b
}

func (c *CalculadoraA) Subtracao(a int, b int) int {
	fmt.Println("CalculadoraA.Subtracao")
	return a - b
}

func (c *CalculadoraA) Multiplicacao(a int, b int) int {
	fmt.Println("CalculadoraA.Multiplicacao")
	return a * b
}

func (c *CalculadoraA) Divisao(a int, b int) int {
	fmt.Println("CalculadoraA.Divisao")
	return a / b
}
