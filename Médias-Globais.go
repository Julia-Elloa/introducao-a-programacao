package main

import "fmt"

func main() {

	int alunos

	fmt.Println("Informe a quantidade de alunos: ")
	fmt.Scan(alunos)

	var nota1 [alunos]float
	var nota2 [alunos]float
	var nota3 [alunos]float
	var media [alunos]float
	var i int

	for i:=0; i <= alunos; i++{
		fmt.Scan(&nota1[i])
		fmt.Scan(&nota2[i])
		fmt.Scan(&nota3[i])

		media[i] = (nota1[i] + nota2[i] + nota3[i])/3
	}

	fmt.Println("Médias globais: ")

	for i:=0; i <= alunos; i++{
		fmt.Println(media[i])
	}

}
