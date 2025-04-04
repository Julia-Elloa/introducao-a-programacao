package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, d float64
  var determinante float64

      fmt.Scan(&a)
      fmt.Scan(&b)
      fmt.Scan(&c)
      fmt.Scan(&d)

    determinante = math.Round((a * d) - (b * c))

    fmt.Println("O VALOR DO DETERMINANTE E = ", determinante)

}
