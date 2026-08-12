package main

import (
	"fmt"
	"time"
)


func exibirMensagem(){
	fmt.Println("Olá de uma Goroutine!")
}

func main() {
	// é um processo, que irá disputar recursos do computador para ser exercutar, por isso do termo concorrencia. 
	go exibirMensagem()
	time.Sleep(1 * time.Second)
	fmt.Println("Olá Main Function")

	// Com o sleep de 1 segundo a goroutine é executada de maneira concorrente.
	// Sem a sleep, a goroutine não é executada, pois a main funcion acaba muito rapido.
}