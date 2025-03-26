programa {

  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    real altura, aresta, volume
    
    leia (altura)
    leia (aresta)

    volume = (3 * aresta * aresta * mat.raiz(3, 2))/2
    volume = (volume * altura)/3

    escreva ("O VOLUME DA PIRAMIDE E = ", mat.arredondar(volume, 2), " METROS CUBICO", "\n")

  }
}
