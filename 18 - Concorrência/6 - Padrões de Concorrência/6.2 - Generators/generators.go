//Generator: basicamente é uma função que vai encapsular uma chamada para uma goroutine, e devolver um canal.
//A vantagem é não chamar ela na função main, basta chamar a função ao qual já faz o encapsulamento da goroutine
//Padrão generator encapsula uma rotina goroutine e retorna um canal de comunicação para ser feita com essa goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	//Variável que recebe o retorno(canal) da função
	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

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
