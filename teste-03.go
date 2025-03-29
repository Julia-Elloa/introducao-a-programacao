package main

import "fmt"

func main() {

	var (
		soma int
		i    int
	)

	for i <= 100 {
		fmt.Println(i)
		soma = soma + i
		i++
	}

	fmt.Print("A soma é: ", soma)
}
