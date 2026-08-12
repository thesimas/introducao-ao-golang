package main

import (
	"fmt"
	"strings"
)

// Constante em GO

const senha string = "SenhaForte123"

func main()  {
		
	// Palavra Reservada : var, nome da variavel, tipo dela, e o valor.

	var texto string = "Primeira variável em GO";
	fmt.Println(texto);

	var number int = 10;
	fmt.Println("Segunda variavel: " , number)

	fmt.Println("Primeira constante: " , senha)

	// Segunda forma de Declarar uma variavel - var, nome da variavel, valor

	var boloDeChocolate = "Bolo de Chocolate"
	fmt.Println("Variavel criada sem declarar o tipo: " , boloDeChocolate);

	// Terceira forma de criar - nome da variavel , e o valor dela. 

	pontoFlutuante := 2.54
	fmt.Println("Variavel declarada de forma curta: " , pontoFlutuante)

	fmt.Println("Variavel String com o Valor Maiusculo: ", strings.ToUpper(boloDeChocolate))
	
}