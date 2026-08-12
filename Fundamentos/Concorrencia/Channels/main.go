package main

import (
	"fmt"
	"time"
)

func main() {

	canalInt := make(chan int)

	// Para escrever um valor em um canal é usando: <-
	canalInt <- 10

	go func() {
		canalInt <- 10
	}()
	time.Sleep(time.Second * 1)

	// É possivel passar o tamanho do canal, passando outro argumento no CHAN
	canalFloat := make(chan float32, 3)
	go func() {
		canalFloat <- 5.5
		canalFloat <- 6.5
		canalFloat <- 10.5
	}()
	time.Sleep(time.Second * 1)

	valor := <-canalFloat
	fmt.Println("Valor do Canal de Pontos Flutuantes: ", valor)

}
