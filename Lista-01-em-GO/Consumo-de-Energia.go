package main

import (
	"fmt"
	"math"
)

func main() {

	var salario float64
	var energia float64
	var custokw float64
	var custototal float64
	var desconto float64

	fmt.Scan(&salario)
	fmt.Scan(&energia)

	custokw = math.Round((salario * 0.7) / 100)
	custototal = math.Round(custokw * energia)
	desconto = math.Round(custototal * 0.9)

	fmt.Println("Custo por KW: R$", custokw)
	fmt.Println("Custo do consumo: R$", custototal)
	fmt.Println("Custo com desconto: R$", desconto)

}
