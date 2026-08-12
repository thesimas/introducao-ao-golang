package main

import (
	"fmt"
	"time"
)

func main() {
	nota := 75

	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	}else if nota >= 70 {
		fmt.Println("Aprovado")
	}else {
		fmt.Println("Reprovado")
	}

	jogadores := map [string]int {
		"Luciano" : 100,
		"João": 50,
		"Lais": 89,
	}

	// Acessar um MAP em go, retorna dois valores, o valor obtido pela chave e um valor booleano, indicando se ele existe ou não. 
	if valor , ok := jogadores["Luciano"]; ok {
		fmt.Println("A pontuação do jogador Luciano é: ",  valor)
	}

	
	fmt.Println("Quando é sabado?")
	
	hoje := time.Now().Weekday();

	switch time.Saturday {
	case hoje + 0:
		fmt.Println("É hoje!")
	case hoje + 1:
		fmt.Println("É amanhã!")
	case hoje + 2:
		fmt.Println("É daqui a 2 dias!")
	case hoje + 3:
		fmt.Println("É daqui a 3 dias!")
	default:
		fmt.Println("Tá longe ainda...")
	}
}