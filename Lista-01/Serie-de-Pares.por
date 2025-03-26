programa {
  funcao inicio() {
    
    inteiro x, y, i

    leia (x)
    leia (y)
    i = 2

    se (x%2 == 0)
    {
      enquanto (i <= y)
      {
       escreva (i, " ")

       i = i + 2
      }
    }
    senao
    {
      escreva ("O PRIMEIRO NUMERO NAO E PAR", "\n")
    }


  }
}
