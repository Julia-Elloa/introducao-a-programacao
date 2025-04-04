package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
    var b float64
    var c float64
    var delta float64

    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

     delta = math.Round(b * b - (4 * a * c))
     fmt.Println("O VALOR DE DELTA E = ", delta)

}
