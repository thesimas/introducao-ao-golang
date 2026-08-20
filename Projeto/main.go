package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Questoes struct {
	Pergunta string
	Opcoes   []string
	Resposta int
}

type Game struct {
	Nome     string
	Pontos   int
	Questoes []Questoes
}

func (game *Game) init(){
	fmt.Println("Olá, seja Bem-Vindo ao jogo de perguntas feito em GO")
	fmt.Println("Qual é o seu apelido?")
	leia := bufio.NewReader(os.Stdin)

	nome, erro := leia.ReadString('\n')

	if erro != nil {
		panic("Erro ao ler a String")
	}

	game.Nome = strings.TrimSpace(nome);
	fmt.Printf("Vamos ao jogo %s!\n", game.Nome)
}

func (game *Game) lendoCSV(){
	arquivo, erro := os.Open("./Projeto/golang_quiz.csv")
	
	if erro != nil {
		panic("Erro ao ler o arquivo CSV!")
	}

	// defer é a ultima linha a ser executada;
	defer arquivo.Close()

	leitor := csv.NewReader(arquivo)

	linhas, erro := leitor.ReadAll()

	if erro != nil {
		panic("Erro ao ler as linhas do CSV!")
	}

	for index, linha := range linhas {
		if index > 0 {
			repostaCorreta, _ := toInt(linha[5]) 
			questao := Questoes{
				Pergunta: linha[0],
				Opcoes: linha[1:5],
				Resposta: repostaCorreta,
			}
			game.Questoes = append(game.Questoes, questao)
		}
	}

}

func toInt(s string) (int , error ){
	inteiro , erro := strconv.Atoi(strings.TrimSpace(s))

	if erro != nil {
		return 0, errors.New("Não é permitido caracetere diferente de número, informe novamente!")
	}

	return inteiro, nil
}

func (game *Game) run(){
	for index, questao := range game.Questoes {
		fmt.Printf("\033[33m %d %s \033[0m\n", index+1, questao.Pergunta);
		
		// Exibindo as opções:
		for indexOpcoes , opcao := range questao.Opcoes {
			fmt.Printf("[%d] - %s\n" , indexOpcoes+1, opcao)
		}
		// Obtendo a reposta do usuario
		fmt.Println("Qual é a resposta correta?")
		var reposta int 
		var erro error

		for {
			leitor := bufio.NewReader(os.Stdin)
			
			leia, _ := leitor.ReadString('\n')

			reposta, erro = toInt(leia[:len(leia)-1])

			if erro != nil {
				fmt.Println(erro.Error())
				continue
			}
		
			break
		}

		if  questao.Resposta == reposta {
			fmt.Println("\033[32m Reposta Correta! \033[0m")
			game.Pontos += 10
		}else {
			fmt.Println("\033[31m Reposta Incorreta! \033[0m")
			
		}
		fmt.Println()
	}
}

func main() {
	jogo := &Game{}
	go jogo.lendoCSV()
	jogo.init()
	jogo.run()

	fmt.Printf("Fim de jogo, %s você fez %d pontos!\n", jogo.Nome, jogo.Pontos)
}
