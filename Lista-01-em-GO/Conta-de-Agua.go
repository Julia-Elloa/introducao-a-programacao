package main

import "fmt"

func main() {

	var conta int
	var agua float32
	var consumidor string
	var valor float32

	fmt.Scan(&conta)
	fmt.Scan(&agua)
	fmt.Scan(&consumidor)

	if consumidor == "R" || consumidor == "r" {
		valor = (agua * 0.05) + 5
	} else if consumidor == "C" || consumidor == "c" {
		valor = 500 + (agua-80)*0.25
	} else if consumidor == "I" || consumidor == "i" {
		valor = 800 + (agua-100)*0.04
	}

	fmt.Println("\n", "CONTA = ", conta)
	fmt.Println("VALOR DA CONTA = ", valor)

}
