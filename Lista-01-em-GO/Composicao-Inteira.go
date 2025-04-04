package main

import "fmt"

func main() {

	var n1 int
    var n2 int
    var n3 int
    var numero int

    fmt.Scan(n1)
    fmt.Scan(n2)
    fmt.Scan(n3)

    if n1 > 10 || n2 > 10 || n3 > 10 {
      fmt.Println("DIGITO INVALIDO")
    } else {
      numero = (n1 * 100) + (n2 * 10) + n3
      fmt.Println(numero, ", ", numero*numero)
    }

}
