package main

import "fmt"

func main() {

	// Array
	var gavetas [2]string
	gavetas[0] = "Copos"
	gavetas[1] = "Pratos"

	// Slice
	var comodo []string
	comodo = append(comodo, "Quarto", "Cozinha", "Sala", "Varanda")
	fmt.Println(comodo)
	fmt.Println(comodo[1:4])


	// HashMaps
	// Map de Pessoas, onde a chave é uma string e o valor é um inteiro.
	var pessoas = map[string]int{}
	pessoas["Luciano"] = 26
	pessoas["João"] = 15

	fmt.Println("O Luciano tem ", pessoas["Luciano"], " anos!")
	fmt.Println(pessoas)

	fmt.Println("João contêm no meu Map?")

	// Idade irá receber o valor de pessoas["João"], a variavel OK irá agir como um valor Booleano
	if idade, ok := pessoas["João"]; ok {
		fmt.Println("João está no map?", ok ,"\nSua idade é: ", idade)
	}else {
		fmt.Println("João não está no map!")
	}


	// Para remover um valor no map se usa a palavra reservado DELETE

	delete(pessoas, "João")

	fmt.Println(pessoas)
}