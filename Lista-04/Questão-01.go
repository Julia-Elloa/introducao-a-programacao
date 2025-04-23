package main

import "fmt"

func main () {
	var num [10]int
	var nulo bool
	
	i := 0

	for i < 10 {
		fmt.Scan(&num[i])
		i++
	}

	i = 0
	for i < 10 {
		nulo = true
		if num[i] > 50 {
			fmt.Println(num[i], " ", i)
		} else {
			nulo = false
		}
		i++
	}

	if nulo == false {
		fmt.Println("Nenhum dos números é superior a 50")
	}
}
