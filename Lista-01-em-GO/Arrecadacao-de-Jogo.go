package main

import "fmt"

func main() {

	nteste := 0

	fmt.Println("Insira a quantidade de jogos: ")
	fmt.Scan(nteste)

	var pessoas [1000]float32
	var popular [1000]float32
	var geral [1000]float32
	var arquibancada [1000]float32
	var cadeiras [1000]float32
	var soma [1000]float32
	i := 0

	for i = 0; i <= nteste; i++ {

		fmt.Println(" ")
		fmt.Println("Informações sobre o jogo ", i+1)
		fmt.Scan(pessoas[i])
		fmt.Scan(popular[i])
		fmt.Scan(geral[i])
		fmt.Scan(arquibancada[i])
		fmt.Scan(cadeiras[i])

		if pessoas[i] > 1 {
			pessoas[i] = pessoas[i] / 100
		}
		if popular[i] > 1 {
			popular[i] = popular[i] / 100
		}
		if geral[i] > 1 {
			geral[i] = geral[i] / 100
		}
		if arquibancada[i] > 1 {
			arquibancada[i] = arquibancada[i] / 100
		}
		if pessoas[i] > 1 {
			cadeiras[i] = cadeiras[i] / 100
		}

		soma[i] = (pessoas[i] * popular[i]) + (pessoas[i] * geral[i] * 5) + (pessoas[i] * arquibancada[i] * 10) + (pessoas[i] * cadeiras[i] * 20)

	}

	i = 0
	fmt.Println(" ")

	for i = 0; i <= nteste; i++ {

		fmt.Println("A RENDA DO JOGO N. ", i+1, " É = ", soma[i])

	}

}
