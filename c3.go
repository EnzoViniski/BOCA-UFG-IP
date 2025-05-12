package main

import "fmt"

func main() {
	var B, E int

	fmt.Scan(&B, &E)

	resultado := B
	final := 0
	if E == 0 {
		final = 1
	} else {
		for i := 1; i < E; i++ {
			resultado = resultado * B
		}
		final = resultado
	}

	fmt.Println(final)
}
