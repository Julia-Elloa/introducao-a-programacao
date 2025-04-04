package main

import (
	"fmt"
	"math"
)

func main() {

	var num, contador int
	var soma, i float64

	fmt.Scan(num)
	soma = 0
	i = 1

	if num > 1 {
		for contador <= num {
			soma = math.Round(soma + (1 / i))
			i = i + 1
		}
	}

	fmt.Println(soma)

}
