package main

import (
	"fmt"
	"math"
)

func main() {
	var peso, altura float64
	fmt.Scan(&peso, &altura)

	imc := peso / math.Pow(altura, 2)

	var classificacao string
	switch {
	case imc < 18.5:
		classificacao = "Abaixo do peso"
	case imc < 25:
		classificacao = "Peso normal"
	case imc < 30:
		classificacao = "Sobrepeso"
	default:
		classificacao = "Obesidade"
	}

	fmt.Println(classificacao)
}
