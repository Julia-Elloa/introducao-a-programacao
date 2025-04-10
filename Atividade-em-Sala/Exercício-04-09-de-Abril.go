package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var l [10]int
	var x int
	
	fmt.Println("Preencha o vetor de", n, " casas:")
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	fmt.Println("Escolha um valor para consultar:")
	fmt.Scan(&x)

	posicao := -1

	for j := 0; j < n; j++ {
		if l[j] == x {
			posicao = j + 1
		}
	}

	fmt.Println("Posição do número", x, "no vetor: ")
	fmt.Println(posicao)

}
