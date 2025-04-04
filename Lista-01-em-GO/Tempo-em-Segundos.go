package main

import "fmt"

func main() {

	var horas, minutos, segundos int

    fmt.Scan(horas)
    fmt.Scan(minutos)
    fmt.Scan(segundos)

    segundos = (horas * 3600) + (minutos * 60) + segundos

    fmt.Println("O TEMPO EM SEGUNDOS E = ", segundos)

}
