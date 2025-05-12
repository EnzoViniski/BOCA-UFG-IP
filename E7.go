package main

import "fmt"

func main (){
	var N int
	
	fmt.Scan(&N)  
	
	vetora := make([]int, N)

	vetorb := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&vetora[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&vetorb[i])
	}

	soma := 0 

    calculo := make([]int, N)
    
	for i := 0; i < N; i++ {
		calculo[i] = vetora[i] * vetorb[i]
		soma += calculo[i]
	}

	fmt.Print(soma)
}

