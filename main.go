package main

import (
	"fmt"

	"github.com/ivocalado/hello-go-interfaces/pkg"
)

func main() {
	chamador := pkg.NewChamador(&pkg.CalculadoraA{})
	fmt.Println(chamador.Soma(1, 2))
	fmt.Println(chamador.Subtracao(1, 2))
	fmt.Println(chamador.Multiplicacao(1, 2))
	fmt.Println(chamador.Divisao(1, 2))

	fmt.Println("-----------------")
	chamador = pkg.NewChamador(&pkg.CalculadoraB{})
	fmt.Println(chamador.Soma(1, 2))
	fmt.Println(chamador.Subtracao(1, 2))
	fmt.Println(chamador.Multiplicacao(1, 2))
	fmt.Println(chamador.Divisao(1, 2))
}
