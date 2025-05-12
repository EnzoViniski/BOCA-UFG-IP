package main

import (
	"fmt"
)

func main() {
	var nota1, nota2, nota3 float64
	var faltas int

	fmt.Scan(&nota1, &nota2, &nota3, &faltas)

	media := (nota1 + nota2 + nota3) / 3

	if faltas > 16 {
		fmt.Println("Reprovado por falta")
		return
	}

	switch {
	case media >= 7:
		fmt.Println("Aprovado")
	case media >= 5 && media < 7:
		fmt.Println("Prova final")
	default:
		fmt.Println("Reprovado")
	}
}
