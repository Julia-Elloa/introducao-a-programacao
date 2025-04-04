package main

import "fmt"

func main() {

	var horas int
    var valor int

    fmt.Scan(&horas)

    if horas%3 == 0 {
      valor = horas * 10
    } else {
      valor = ((horas - horas%3)/3) * 10 + (horas%3) * 5
    }

    fmt.Println("O VALOR A PAGAR E = ", valor)

}
