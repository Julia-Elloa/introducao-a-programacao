package main

import "fmt"

func main() {

	var salario float64

    fmt.Scan(&salario)

    if salario <= 300 {
      salario = salario * 1.5
    } else {
      salario = salario * 1.3
    }

    fmt.Println("SALARIO COM REAJUSTE = ", salario)

}
