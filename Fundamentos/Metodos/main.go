package main

import "fmt"

type Endereco struct {
	Rua    string
	Cep    string
	Cidade string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

func (pessoa Pessoa) Apresentar() {
	fmt.Println("Irei apresentar essa pessoa!")
	fmt.Println("Nome: ", pessoa.Nome)
	fmt.Println("Idade: ", pessoa.Idade)
	fmt.Println("Email: ", pessoa.Email)
	fmt.Println("Rua: ", pessoa.Endereco.Rua)
	fmt.Println("Cep: ", pessoa.Endereco.Cep)
	fmt.Println("Cidade", pessoa.Endereco.Cidade)

}

func main() {
	pessoa := Pessoa{
		Nome: "Luciano",
		Idade: 26,
		Email: "luciano@gmail.com",
		Endereco: Endereco {
			Cep: "0000-0000", 
			Rua: "Rua das Flores",
			Cidade: "Florianopolis",
		},
	}
	
	pessoa.Apresentar()
}
