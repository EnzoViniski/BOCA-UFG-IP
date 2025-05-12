package main

import "fmt"

func main (){
	var N int
	
	fmt.Scan(&N)

	vetor := make([]float64, N)
    total := 0.0 

	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
		total += vetor[i]
	}

    var N2 float64 = float64(N)
    
	media := total / N2 
    
	for i := 0; i < N; i++ {
		if vetor[i] > media {
			fmt.Printf("%.2f\n", vetor[i])
		}
	}
}

