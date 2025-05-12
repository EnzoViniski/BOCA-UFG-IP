package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	numeros := []int{a, b, c}

	sort.Ints(numeros)

	fmt.Println(numeros[0], numeros[1], numeros[2])
}
