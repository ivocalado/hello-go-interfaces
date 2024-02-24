package pkg

type CalculadoraInterface interface {
	Soma(a int, b int) int
	Subtracao(a int, b int) int
	Multiplicacao(a int, b int) int
	Divisao(a int, b int) int
}
