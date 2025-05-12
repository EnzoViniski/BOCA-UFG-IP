package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	// Coloca os lados em uma slice e ordena para facilitar a verificação
	lados := []int{x, y, z}
	sort.Ints(lados)

	// Verifica se pode formar triângulo
	if lados[0]+lados[1] <= lados[2] {
		fmt.Println("Nao forma triangulo")
		return
	}

	// Classifica o triângulo
	if lados[0] == lados[1] && lados[1] == lados[2] {
		fmt.Println("Equilatero")
	} else if lados[0] == lados[1] || lados[1] == lados[2] || lados[0] == lados[2] {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Escaleno")
	}
}
