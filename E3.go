package main

import "fmt"

func main (){
	var N, n2 int
	
	fmt.Scan(&N)
    
	if N == 1 {
		fmt.Scan(&n2)
		fmt.Print("0")
	} else {
	vetor := make([]int, N)
	
	N2 := N - 1 
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	vetor2 := make([]int, N)

	N3 := N2 - 1 
	vetor2[0] = vetor [1]

	fmt.Print( vetor2[0])
	for j := 1; j < N2; j++ {
		j2 := j - 1
		j3 := j + 1
		vetor2[j] = vetor[j2] + vetor[j3]
		fmt.Print(" ",vetor2[j])
	} 

	vetor2[N2] = vetor[N3]

	fmt.Print(" ", vetor2[N2])
	}
}

