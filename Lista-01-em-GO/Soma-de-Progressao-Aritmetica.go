package main

import "fmt"

func main() {

	var inicio, razao, elementos, soma int

	fmt.Scan(inicio)
	fmt.Scan(razao)
	fmt.Scan(elementos)
	soma = inicio
	i := 1

	for i < elementos {
		inicio = inicio + razao
		soma = soma + inicio
		i = i + 1
	}

	fmt.Println(soma)

}
