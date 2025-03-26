programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {

    real a, b, c, d
    real determinante

      leia (a)
      leia (b)
      leia (c)
      leia (d)

    determinante = (a * d) - (b * c)

    escreva ("O VALOR DO DETERMINANTE E = ", mat.arredondar(determinante, 2))
    
  }
}
