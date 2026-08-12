package main

import "fmt"

type Pessoa struct {
	Nome string
}


// Ponteiros servem para apontar o endereço de memória de uma variavel.
// Assim você consegue obter o valor da variavel e o endereço dela.

// var x int = 100, essa variavel está no endereço > 0x0201 com o valor 100
// var y *int = &x, essa variavel é o nosso ponteiro, ta no endereço de memória > 0x0206 com o valor 0x0201, 
// que é o endereço de memória da variavel x. 
// Para mostrar o valor de x na variavel y, quando for printar essa variável, se remove o & (e comercial) da atribuição. 
// Ponteiro é chamado também de passar o valor por Referência, assim é possivel mudar o valor da variavel original.

func main() {
	var pessoa1 Pessoa = Pessoa{Nome: "Luciano"}
	// var pessoa2 Pessoa = Pessoa{Nome: "João"}
	fmt.Println("Nome da Pessoa 1: ", pessoa1.Nome)

	var pessoa3 *Pessoa = &pessoa1

	fmt.Println("Endereço de Memória da Pessoa 1: ", &pessoa1.Nome)

	fmt.Println("Nome da Pessoa 3, que aponta para a pessoa 1: ", pessoa3.Nome)
	// fmt.Println(pessoa2)

	pessoa3.Nome = "Francisco"

	fmt.Println("Nome da Pessoa 3, reatribuido: ", pessoa3.Nome)
	fmt.Println("Nome da Pessoa 1: ", pessoa1.Nome)

}