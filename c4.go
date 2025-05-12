package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	for n < 0 || n > 0 {
		final := 0

		for i := 1; i < n; i++ {
			test := i * i
			if test == n {
				fmt.Println(n, "eh quadrado perfeito")
				final = 1
			}
		}

		if final == 0 {
			fmt.Println(n, "nao eh quadrado perfeito")
		}
		fmt.Scan(&n)
	}

}
