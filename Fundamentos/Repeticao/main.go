package main

import "fmt"

func main() {

	// for inicializacao; condicao; fim da interação
	soma := 0
	for i := 0; i < 10; i++ {
		soma += i
	}
	fmt.Println("Soma: ", soma)

	// Simulando WHILE em GO
	soma = 0
	for soma < 20 {
		fmt.Println(((soma /5) +1), "ª interação!")
		soma += 5
	}
	fmt.Println("Soma: ", soma)

	// Pecorrendo listas
	gavetas := []string {"Copo", "Talher", "Prato"}
	for i := 0; i < len(gavetas); i ++ {
		fmt.Println((i+1), "º Elemento: ", gavetas[i])
	}

	// Range
	numeros := []int {2, 4, 5, 7, 0}
	for chave, valor := range numeros {
		fmt.Println("Chave: ", chave, " Valor: ", valor)
	}

	usuarios := map[string] string {
		"nome": "João" ,
		"idade": "35",
	}
	for chave, valor := range usuarios {
		fmt.Println("Chave: ", chave, " Valor: ", valor)
	}
}