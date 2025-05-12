package main

import (
	"fmt"
)

func main() {
	var vendasBrutas float64
	fmt.Scan(&vendasBrutas)

	const salarioBase = 500.00
	const comissaoPercentual = 0.09
	const bonusLimite = 15000.00
	const bonusValor = 800.00

	comissao := vendasBrutas * comissaoPercentual

	salarioTotal := salarioBase + comissao

	if vendasBrutas > bonusLimite {
		salarioTotal += bonusValor
	}

	fmt.Printf("%.5f\n", salarioTotal)
}
