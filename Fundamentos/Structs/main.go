package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua string
	Cep string
	Cidade string
}

func main() {
	cliente1 := Cliente{
		Nome:     "Luciano",
		Idade:    26,
		Endereco: Endereco{
			Rua: "Rua das flores",
			Cep: "0000-0000",
			Cidade: "Floarinópolis",
		},
	}

	fmt.Println(cliente1)
	fmt.Println("Faltou o Email!")

	cliente1.Email = "luciano@gmail.com"

	fmt.Println(cliente1)
}