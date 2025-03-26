programa {
  funcao inicio() {
    
    inteiro horas
    inteiro valor

    leia (horas)

    se (horas%3 == 0)
    {
      valor = horas * 10
    }
    senao
    {
      valor = ((horas - horas%3)/3) * 10 + (horas%3) * 5

    }

    escreva ("O VALOR A PAGAR E = ", valor, "\n")


  }
}
