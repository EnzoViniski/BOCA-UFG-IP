package main

import "fmt"

func main (){
	var N int
	
	fmt.Scan(&N)  
	
	vetor := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	maior, atual := 1, 1
     
	for i := 1; i < N; i++ {
		if vetor[i] > vetor[i - 1]{
			atual ++
			if atual > maior {
				maior = atual
			}
		} else {
			atual = 1
		}
		
	}


	fmt.Print(maior)
}

