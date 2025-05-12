package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	vetor := make([]int, N)
	vetor2 := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	j := 0
	
	N2 := N - 1
	for i := N2 ; i >= 0; i-- {
		
		vetor2[j] = vetor[i]
		j += 1
	}

    AF := 0
	for i := 0; i < N; i++ {
		if vetor[i] != vetor2[i]{
			AF = 1
		}

	}

	if AF == 1 {
		fmt.Print("NAO")
	} else {
		fmt.Print("SIM")
	}
}
