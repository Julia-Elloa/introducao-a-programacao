package main

import "fmt"

func main() {

	var n int
    	var i int

    i = 2

    fmt.Scan(&n)

    if 5 < n && n < 2000 {
        for i <= n {
		fmt.Println(i, "^", i, " = ", i * i)
		i = i + 2
     	 }
    } else {
      fmt.Println("O número deve ser maior que 5 e menor que 2000")
    }

}
