package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	// Corrigir problemas de ponto flutuante
	for a := 0.0; a <= x+1e-9; a += 0.1 {
		sinA := a -
			(math.Pow(a, 3) / float64(3*2*1)) +
			(math.Pow(a, 5) / float64(5*4*3*2*1)) -
			(math.Pow(a, 7) / float64(7*6*5*4*3*2*1))

		fmt.Printf("%.1f %.4f\n", a, sinA)
	}
}
