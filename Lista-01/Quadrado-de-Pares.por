programa {
  funcao inicio() {

    inteiro n
    inteiro i

    i = 2

    leia (n)

    se (5 < n e n < 2000)
    {
        enquanto (i <= n)
      {
        escreva (i, "^", i, " = ", i * i, "\n")
        i = i + 2
      }
    }
    senao
    {
      escreva ("O número deve ser maior que 5 e menor que 2000", "\n")
    }
    
  }
}
