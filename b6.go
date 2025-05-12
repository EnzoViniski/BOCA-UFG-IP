package main

import (
	"fmt"
)

func main() {
	var conta int
	var limite, saldoInicial, depositos, retiradas float64

	fmt.Scan(&conta, &limite, &saldoInicial, &depositos, &retiradas)

	saldoFinal := saldoInicial + depositos - retiradas

	if saldoFinal > limite {
		fmt.Printf("Credito excedido: %.5f\n", saldoFinal)
	} else {
		fmt.Printf("Saldo: %.5f\n", saldoFinal)
	}
}
