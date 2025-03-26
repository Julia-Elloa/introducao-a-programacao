programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {

    real raio
    real altura
    real custo
    
    leia (raio)
    leia (altura)

    custo = 2 * (3.14159 * raio * raio) + (2 * 3.14159 * raio * altura) * 100
    escreva ("O VALOR DO CUSTO E = ", mat.arredondar(custo, 2))
    
  }
}
