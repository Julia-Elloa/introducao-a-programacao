package main

import "fmt"

func main() {

	var x, y int

    fmt.Scan(x)
    fmt.Scan(y)
    i := 2

    if x%2 == 0 {
      for i <= y {
       fmt.Print(i, " ")
       i = i + 2
      }
    } else {
      fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
    }

}
