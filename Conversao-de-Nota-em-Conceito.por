programa {
  funcao inicio() {
    
    real nota

    leia (nota)

    se (nota > 9)
    {
      escreva ("NOTA = ", nota, " CONCEITO = A", "\n")
    }
    senao se (nota > 7.5)
    {
      escreva ("NOTA = ", nota, " CONCEITO = B", "\n")
    }
    senao se (nota > 6)
    {
      escreva ("NOTA = ", nota, " CONCEITO = C", "\n")
    }
    senao
    {
      escreva ("NOTA = ", nota, " CONCEITO = D", "\n")
    }

  }
}
