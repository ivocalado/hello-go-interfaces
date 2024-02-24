package pkg

type Chamador struct {
	calculadora CalculadoraInterface
}

func NewChamador(calculadora CalculadoraInterface) *Chamador {
	return &Chamador{calculadora: calculadora}
}

func (c *Chamador) Soma(a int, b int) int {
	return c.calculadora.Soma(a, b)
}

func (c *Chamador) Subtracao(a int, b int) int {
	return c.calculadora.Subtracao(a, b)
}

func (c *Chamador) Multiplicacao(a int, b int) int {
	return c.calculadora.Multiplicacao(a, b)
}

func (c *Chamador) Divisao(a int, b int) int {
	return c.calculadora.Divisao(a, b)
}
