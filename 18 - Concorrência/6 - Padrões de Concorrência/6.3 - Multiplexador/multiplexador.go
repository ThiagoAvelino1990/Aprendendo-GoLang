//Multiplexador serve para pegar os dados de dois canais e retornar em um canal só
package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	//for para impressão dos dados
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

//função que recebe dois canais do tipo de string e retorna um canal. Todos como canais que só recebem valores
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	//Criando um canal
	canalDeSaida := make(chan string)

	//goroutine em uma função anônima
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

//função que recebe string como parâmetro e retorna um canal do tipo string, neste caso o canal é de uma direção apenas, só recebe os dados
func escrever(texto string) <-chan string {
	//Declaração do canal
	canal := make(chan string)
	//Função anônima com goroutine
	go func() {
		//For infinito aonde o canal recebe o valor passado na variável texto
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
