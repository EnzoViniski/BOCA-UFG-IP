package main 

import (
	"fmt"
	"sort"
)

func main (){
	var N int 

	fmt.Scan(&N)
	vetor := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}


	sort.Ints(vetor)
	for i, num := range vetor {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()

}