package main

import "fmt"

func main () {
	var num1 [10]int
	var num2 [5]int
	i := 0
	soma := 0

	for i < 10 {
		fmt.Scan(&num1[i])
		i++
	}

	i = 0

	for i < 5 {
		fmt.Scan(&num2[i])
		soma = soma + num2[i]
		i++
	}

	var result1 [10]int
	var result2 [10]int

	i = 0
	par := 0
	impar := 0

	for i < 10 {
		if num1[i]%2 == 0 {
			result1[par] = num1[i] + soma
			par++
		} else {
			result2[impar] = num1[i] + soma
			impar++
		}
		i++
	}

	i = 0
	fmt.Print("Primeiro vetor resultante [ ")
	for i < par {
		fmt.Print(result1[i], " ")
		i++
	}
	fmt.Println("]")

	i = 0
	fmt.Print("Segundo vetor resultante [ ")
	for i < impar {
		fmt.Print(result2[i], " ")
		i++
	}
	fmt.Println("]")
}
