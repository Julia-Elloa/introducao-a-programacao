package main

import "fmt"

func main() {

	var dividendo, divisor float64

	fmt.Println("Insira o dividendo:")
	fmt.Scan(&dividendo)
	fmt.Println("Insira o divisor:")
	fmt.Scan(&divisor)

	e, r := divisao(dividendo, divisor)

	if !e {
		fmt.Println("Resultado da divisão: ", r)
	} else {
		fmt.Println("Divisão não permitida")
	}
}

func divisao(dividendo, divisor float64) (bool, float64) {

	erro := false
	resultado := 0.0

	if divisor == 0 {
		erro = true
	} else {
		resultado = dividendo/divisor
	}

	return erro, resultado
}
