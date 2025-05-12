package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	fmt.Scan(&X)

	result := 0.0
	sinal := -1.0
	mod := 1.0
	E := 1.0

	for n := 2; n < 21; n++ {
		sinal = sinal * (-1)
		mod = mod * float64(n)
		E = E + 1.0
		result = result + (sinal)*math.Pow(X, E)/mod

	}

	fmt.Printf("%.5f\n", result)
}
