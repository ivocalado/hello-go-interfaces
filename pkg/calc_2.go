package pkg

import "fmt"

type CalculadoraB struct {
}

var _ CalculadoraInterface = (*CalculadoraB)(nil)

func (c *CalculadoraB) Soma(a int, b int) int {
	fmt.Println("CalculadoraB.Soma")
	return a + b
}

func (c *CalculadoraB) Subtracao(a int, b int) int {
	fmt.Println("CalculadoraB.Subtracao")
	return a - b
}

func (c *CalculadoraB) Multiplicacao(a int, b int) int {
	fmt.Println("CalculadoraB.Multiplicacao")
	return a * b
}

func (c *CalculadoraB) Divisao(a int, b int) int {
	fmt.Println("CalculadoraB.Divisao")
	return a / b
}
