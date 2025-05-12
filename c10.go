package main

import (
	"fmt"
)

func main() {
	var pinicial, despesa, deltaP, pmin float64
	var qinicial, deltaQ int

	fmt.Scan(&pinicial, &qinicial, &despesa, &deltaP, &deltaQ, &pmin)

	fmt.Printf("%5s %20s %8s\n", "Preco", "Ingressos Vendidos", "Lucro")

	lucroMax := -1.0
	precoMaxLucro := 0.0
	qtdMaxLucro := 0

	preco := pinicial
	qtd := qinicial

	for preco >= pmin {
		lucro := (preco * float64(qtd)) - despesa

		fmt.Printf("%5.2f %13d %16.2f\n", preco, qtd, lucro)

		if lucro > lucroMax {
			lucroMax = lucro
			precoMaxLucro = preco
			qtdMaxLucro = qtd
		}

		preco -= deltaP
		qtd += deltaQ
	}

	fmt.Printf("Lucro maximo: %.2f\n", lucroMax)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.\n", precoMaxLucro, qtdMaxLucro)
}
