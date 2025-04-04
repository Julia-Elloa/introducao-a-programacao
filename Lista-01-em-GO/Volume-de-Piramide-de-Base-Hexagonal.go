package main

import (
	"fmt"
	"math"
)

func main() {

	var altura, aresta, volume float64
    
    fmt.Scan(altura)
    fmt.Scan(aresta)

    volume = math.Round((3 * aresta * aresta * math.Sqrt(3))/2)
    volume = math.Round((volume * altura)/3)

    fmt.Println("O VOLUME DA PIRAMIDE E = ", volume, " METROS CUBICO")

}
