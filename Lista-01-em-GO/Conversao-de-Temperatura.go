package main

import (
	"fmt"
	"math"
)

func main() {

	var nlinhas int

	fmt.Scan(nlinhas)

	var f [1000]float64
	var c [1000]float64
	var i int

	i = 0

	for i < nlinhas {
		fmt.Scan(f[i])

		c[i] = math.Round((f[i] - 32) / 1.8)

		i = i + 1
	}

	i = 0

	for i < nlinhas {
		fmt.Println(f[i], "FAHRENHEIT EQUIVALE A ", c[i], " CELSIUS")
		i = i + 1
	}

}
