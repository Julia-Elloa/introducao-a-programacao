programa {

  inclua biblioteca Matematica --> mat
  funcao inicio()
  
  {
    
    real n1, n2, n3, media

    leia (n1) leia (n2) leia (n3)

    media = (n1 + n2 +  n3)/3

    se (media >= 6)
    {
      escreva ("MEDIA = ", mat.arredondar(media, 2), "\n")
      escreva ("APROVADO")
    }
    senao
    {
      escreva ("MEDIA = ", mat.arredondar(media, 2), "\n")
      escreva ("REPROVADO")
    }

  }
}
