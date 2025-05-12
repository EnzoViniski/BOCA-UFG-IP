package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	cosX := 1.0
	sinal := -1.0
	potencia := 2
	fat := 1.0

	for i := 2; i <= potencia; i++ {
		fat *= float64(i)
	}

	for i := 1; i < 20; i++ {
		termo := (sinal * math.Pow(x, float64(potencia))) / fat
		cosX += termo

		sinal *= -1
		potencia += 2

		fat *= float64(potencia-1) * float64(potencia)
	}

	valor := math.Cos(x)
	dif := cosX - valor

	fmt.Printf("%.4f %.4f %.4f\n", cosX, valor, dif)
}
