package main

import (
	"fmt"
	"time"
)

func producer(c chan int){
	for x := 0; x < 5; x++ {
		c <- x
	}
	close(c)
}

func consumer(c chan int){
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Consumer finalizado")
}

func main() {

	canalInt := make(chan int)

	// Para escrever um valor em um canal é usando: <-
	// canalInt <- 10

	go func() {
		canalInt <- 10
		close(canalInt)
	}()
	time.Sleep(time.Second * 1)
	valorInt := <-canalInt
	fmt.Println("Valor lido do canalInt: ", valorInt)

	// É possivel passar o tamanho do canal, passando outro argumento no CHAN
	canalFloat := make(chan float32, 3)

	go func() {
		canalFloat <- 5.5
		canalFloat <- 6.5
		canalFloat <- 10.5
		close(canalFloat)
	}()
	time.Sleep(time.Second * 1)

	valor := <-canalFloat
	fmt.Println("Valor do Canal de Pontos Flutuantes: ", valor)

	// Para ler o terceiro valor do canal da go routine, eu preciso esvazia-lo
	<- canalFloat
	valor = <- canalFloat
	fmt.Println("Terceiro valor do Canal de Pontos Flutuantes: ", valor)
	
	canalProdutor := make(chan int);
	go producer(canalProdutor)
	go consumer(canalProdutor)
	go consumer(canalProdutor)

	time.Sleep(time.Second * 1)
}
