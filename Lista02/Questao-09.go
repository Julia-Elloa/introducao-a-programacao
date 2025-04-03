package main

import "fmt"

func main() {

	var compra float64
	var venda float64

	fmt.Scan(compra)

    if compra < 0 {
    	if compra < 50 {
    		venda = compra * 1.3
    	} else if compra >= 30 {
    		venda = compra * 1.4
    	} else if compra >= 10 {
    		venda = compra * 1.5
    	} else if compra > 0 {
    		venda = compra * 1.7
    	}
    	fmt.Println(venda)
    } else {
        fmt.Println("Valor inválido")
    }
    

}
