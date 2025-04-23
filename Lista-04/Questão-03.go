package main

import "fmt"

func main () {
	var num [10]int
	var somapar, impar int
	
	i := 0

	for i < 10 {
		fmt.Scan(&num[i])
		i++
	}

	i = 0
	fmt.Print("Números pares digitados: ")
	for i < 10 {
		if num[i]%2 == 0 {
			fmt.Print(num[i], " ")
			somapar += num[i]
		}
		i++
	}
	fmt.Println("")

	fmt.Println("Soma dos números pares digitados:", somapar)

	i = 0
	fmt.Print("Números impares digitados: ")
	for i < 10 {
		if num[i]%2 != 0 {
			fmt.Print(num[i], " ")
			impar++
		}
		i++
	}
	fmt.Println("")

	fmt.Println("Quantidade de números ímpares digitados:", impar)
}
