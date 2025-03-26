programa {
  funcao inicio() {
    
    inteiro n1
    inteiro n2
    inteiro n3
    inteiro numero

    leia (n1)
    leia (n2)
    leia (n3)

    se (n1 > 10 ou n2 > 10 ou n3 > 10)
    {
      escreva ("DIGITO INVALIDO")
    }
    senao
    {
      numero = (n1 * 100) + (n2 * 10) + n3
      escreva (numero, ", ", numero*numero)
    }

  }
}
