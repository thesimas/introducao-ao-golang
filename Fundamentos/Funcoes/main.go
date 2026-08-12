package main

import "fmt"

func main() {
	fmt.Println(soma(7, 8))

	var fixo = 5
	multiplica := func(x int) int {
		return x * fixo 
	}

	fmt.Println(multiplica(7))
}

// Letra minuscula a função fica como private
func soma(nota1, nota2 int) int {
	return (nota1 + nota2)
}
