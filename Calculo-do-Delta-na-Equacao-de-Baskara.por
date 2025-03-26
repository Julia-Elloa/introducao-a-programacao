programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    real a 
    real b 
    real c 
    real delta

    leia (a)
    leia (b)
    leia (c)

     delta = b * b - (4 * a * c)
     escreva ("O VALOR DE DELTA E = ", mat.arredondar(delta, 2))

  }
}
