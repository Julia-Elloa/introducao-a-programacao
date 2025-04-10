package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var l [100]int
	var x int

	i := 0
	j := 0

	fmt.Println("Preencha o vetor de", n, " casas:")
	for i < n {
		fmt.Scan(&l[i])
		i++
	}

	fmt.Println("Escolha um valor para consultar:")
	fmt.Scan(&x)

	posicao := -1

	for j < n {
		if l[j] == x {
			posicao = j + 1
		}
		j++
	}

	fmt.Println("Posição do número", x, "no vetor: ")
	fmt.Println(posicao)

}
