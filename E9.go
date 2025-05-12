package main 

import (
	"fmt"
)

func main (){
	var N int 

	fmt.Scan(&N)

	vetor := make([]int, N)

	for i := 0; i < N ; i++ {
		fmt.Scan(&vetor[i])
	}

	vetor2 := make([]int, N)

	soma := 0 

	for j := 0; j < N; j++{
		for i := 0; i <= j ; i++ {
			soma += vetor[i]
		}
		vetor2[j] = soma
		soma = 0
	}

	fmt.Print(vetor2[0])
	for j := 1; j < N; j++{
		fmt.Print(" ",vetor2[j])
	}


}