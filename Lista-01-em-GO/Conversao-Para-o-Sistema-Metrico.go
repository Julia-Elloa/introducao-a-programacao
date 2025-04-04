package main

import (
	"fmt"
	"math"
)

func main() {

	var temperatura float64
    var chuva float64

    fmt.Scan(temperatura)
    fmt.Scan(chuva)

    temperatura = math.Round((temperatura - 32)/1.8)
    chuva = math.Round(chuva * 25.4)

    fmt.Println("O VALOR EM CELSIUS = ")
    fmt.Println("A QUANTIDADE DE CHUVA E = ")

}
