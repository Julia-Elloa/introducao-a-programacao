package main

import "fmt"

func main() {

	var n1, n2, n3, media float32

    	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

    media = (n1 + n2 +  n3)/3

    if media >= 6 {
      fmt.Println("MEDIA = ", media)
      fmt.Println("APROVADO")
    } else {
      fmt.Println("MEDIA = ", media)
      fmt.Println("REPROVADO")
    }

}
