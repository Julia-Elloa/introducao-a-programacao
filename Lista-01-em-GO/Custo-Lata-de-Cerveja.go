package main

import (
	"fmt"
	"math"
)

func main() {

	var raio float64
    var altura float64
    var custo float64
	const pi float64 = 3.14159
    
    fmt.Scan(raio)
    fmt.Scan(altura)

    custo = math.Round(2 * (pi * raio * raio) + (2 * 3.14159 * raio * altura) * 100)
    fmt.Println("O VALOR DO CUSTO E = ", custo)

}
