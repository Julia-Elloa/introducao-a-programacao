package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Escreva um número: ")
	fmt.Scan(numero)

	if numero > 0 {
		fmt.Println("Positivo")
	} else if numero < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Nulo")
	}
}
