package main

import "fmt"

func main() {

	var a, b, s int

	fmt.Println("Insira o primeiro número:")
	fmt.Scan(&a)
	fmt.Println("Insira o segundo número:")
	fmt.Scan(&b)

	s = soma(a, b)

	fmt.Println("Resultado:")
	fmt.Println(s)

}

func soma(a, b int) int {

	return a + b

}
