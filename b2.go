package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N%2 == 0 {
		fmt.Printf("%d Par\n", N)
	} else {
		fmt.Printf("%d Impar\n", N)
	}
}
