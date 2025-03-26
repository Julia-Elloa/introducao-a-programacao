programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    real temperatura
    real chuva

    leia (temperatura)
    leia (chuva)

    temperatura = (temperatura - 32)/1.8
    chuva = chuva * 25.4

    escreva ("O VALOR EM CELSIUS = ", mat.arredondar(temperatura, 2), "\n")
    escreva ("A QUANTIDADE DE CHUVA E = ", mat.arredondar(chuva, 2))
  }
}
