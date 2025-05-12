package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	valores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
	}

	distintos := make(map[int]bool)
	for _, v := range valores {
		distintos[v] = true
	}

	fmt.Println(len(distintos))
}
