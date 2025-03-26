programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    inteiro n, i
    real soma

    leia (n)
    soma = 0
    i = 1

    se (n > 1)
    {
      enquanto (i <= n)
      {
        soma = soma + (1/i)
        i = i + 1
      }
    }

    escreva (mat.arredondar(soma, 6), "\n")

  }
}
