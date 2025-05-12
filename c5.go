package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	test := 0
	for i := 1; i < n; i++ {
		teste := i * (i + 1) * (i + 2)

		if teste == n {
			fmt.Println(n, "eh triangular")
			test = 1
		}
	}
	if test == 0 {
		fmt.Println(n, "nao eh triangular")
	}

}
