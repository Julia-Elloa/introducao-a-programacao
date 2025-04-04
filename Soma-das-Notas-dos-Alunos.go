package main

import "fmt"

func main() {

	var alunos int
	var nota [100]float64
	var soma float64
	i := 0

	fmt.Println("Informe a quantidade de alunos: ")
	fmt.Scan(&alunos)
	fmt.Println("Insira as notas: ")

	for i = 0; i < alunos; i++ {
		fmt.Scan(&nota[i])
		soma = soma + nota[i]
	}

	fmt.Println("A soma das notas é: ", soma)

}
