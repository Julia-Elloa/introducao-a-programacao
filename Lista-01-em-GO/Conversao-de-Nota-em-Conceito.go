package main

import "fmt"

func main() {

	var nota float64

    fmt.Scan(nota)

    if nota > 9 {
      fmt.Println ("NOTA = ", nota, " CONCEITO = A")
    } else if nota > 7.5 {
      fmt.Println ("NOTA = ", nota, " CONCEITO = B")
    } else if nota > 6 {
      fmt.Println ("NOTA = ", nota, " CONCEITO = C")
    } else {
      fmt.Println ("NOTA = ", nota, " CONCEITO = D")
    }

}
