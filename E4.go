package main 

import ("fmt";"math")

func main (){
	var N int 

	fmt.Scan(&N)
	vetor := make([]float64, N)

	for i := 0 ; i < N ; i++ {
		fmt.Scan(&vetor[i])
	}

	j := N - 1
	S := 0.0

	for i := 0 ; i < N/2 ; i++ {
		S += math.Pow(vetor[i] - vetor[j], 3)
		j = j - 1
    }

	fmt.Println(S)
}