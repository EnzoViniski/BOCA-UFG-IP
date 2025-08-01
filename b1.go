package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	if A == 0 {
		fmt.Println("Nao e equacao do segundo grau")
		return
	}

	delta := B*B - 4*A*C

	if delta < 0 {
		fmt.Println("Nao ha raizes reais")
	} else {
		raizDelta := math.Sqrt(delta)
		x1 := (-B + raizDelta) / (2 * A)
		x2 := (-B - raizDelta) / (2 * A)
		fmt.Printf("%.2f %.2f\n", x1, x2)
	}
}
