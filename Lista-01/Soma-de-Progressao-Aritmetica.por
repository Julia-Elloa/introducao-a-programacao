programa {
  funcao inicio() {
    
    inteiro inicio, razao, elementos, soma, i

    leia (inicio)
    leia (razao)
    leia (elementos)
    soma = inicio
    i = 1

    enquanto (i < elementos)
    {
      inicio = inicio + razao
      soma = soma + inicio
      i = i + 1
    }

    escreva (soma, "\n")

  }
}
