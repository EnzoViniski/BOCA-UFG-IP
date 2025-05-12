package main

import "fmt"

func main (){
	var N int
	
	fmt.Scan(&N)

	vetor := make([]float64, N)
	
	fmt.Scan(&vetor[0])
	Xmax := vetor[0]
    Xmin := vetor[0]
	
	for i := 1; i < N; i++ {
		fmt.Scan(&vetor[i])
		
		if vetor[i] > Xmax {
			Xmax = vetor[i]
		}
		if vetor[i] < Xmin {
			Xmin = vetor[i]
		}
	}

	if Xmax == Xmin {

		fmt.Print("0.00")

		for i := 1; i < N; i++ {
			fmt.Print(" 0.00")
		}
	} else {

	Xnormal := (vetor[0] - Xmin )/( Xmax - Xmin )
	fmt.Print("0.00")
	
	for i := 1; i < N; i++ {
       Xnormal = (vetor[i] - Xmin )/( Xmax - Xmin )
	   fmt.Printf(" %.2f",Xnormal)
	}}

}

