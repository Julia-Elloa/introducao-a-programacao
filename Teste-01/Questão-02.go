package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Escreva um número: ")
	fmt.Scan(numero)

	if numero > 20 && numero < 90 {
		fmt.Println("O número está compreendido entre 20 e 90")
	} else {
		fmt.Println("O número não está compreendido entre 20 e 90")
	}
}
