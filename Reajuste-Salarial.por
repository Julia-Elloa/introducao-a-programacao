programa {
  funcao inicio() {
    
    real salario

    leia (salario)

    se (salario < 300)
    {
      salario = salario * 1.5
    }
    senao
    {
      salario = salario * 1.3
    }

    escreva ("SALARIO COM REAJUSTE = ", salario)

  }
}
