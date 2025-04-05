package main

import (
	"fmt"
	"math"
)

func main () {

	var a, b, c, delta float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta = b*b - (4 * a * c)

	if a == 0 {
		fmt.Println("Nao e equacao do segundo grau")
	} else {
		if delta >= 0 {
			var raiz1, raiz2 float64
	
			raiz1 = (math.Sqrt(delta) - b)/(2*a)
			raiz2 = (- math.Sqrt(delta) - b)/(2*a)
	
			fmt.Print(raiz1, " ", raiz2)
		} else if delta < 0 {
			fmt.Println("Nao ha raizes reais")
		}
	}
	
}