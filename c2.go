package main

import "fmt"

func main() {
	var r, l int

	fmt.Scan(&l, &r)

	soma := 0
	tot := 0
	if l < r {
		for i := l; i <= r; i++ {
			div := i % 2
			if div == 0 {
				soma = soma + i
				tot = tot + 1
			}

		}
		med := soma / tot
		fmt.Println(soma)
		fmt.Println(med)
	} else {
		fmt.Println("0")
		fmt.Println("0")
	}
}
