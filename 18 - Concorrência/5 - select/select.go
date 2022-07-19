package main

import (
	"fmt"
	"time"
)

func main() {
	//Declarando canais
	canal1, canal2 := make(chan string), make(chan string)

	//Função anônima com goroutine
	go func() {
		//Loop infinito
		for {
			//Execução em 500 milisegundos.
			time.Sleep(time.Millisecond * 500)
			//Após agaurdar os 500 milisegundos o Canal irá receber o valor
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			//Execução em 2 segundos
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	func() {
		for {
			//com o select o programa ganha tempo de execução
			select {
			//Canal passando o valor para a variável. Irá executar a mensagem a após a conclusão irá para o próximo CASE
			case mensagemCanal1 := <-canal1:
				fmt.Println(mensagemCanal1)
			case mensagemCanal2 := <-canal2:
				fmt.Println(mensagemCanal2)
			}
		}
	}()
}
