//Canal de comunicação. É um canal que pode tanto enviar quanto recever dados.
//o canal tem duas operações básicas e importantes: enviar um dado e receber um dado. Essas operações bloqueam a execução do programa
//CUIDADO COM DEADLOCK - Deadlock é quando você não tem mais nnehum lugar para enviar dado ao canal, mas o canal ainda está esperando um dado.
//Deadlock não é pego em compilação, apenas em tempo de execução.
package main

import (
	"fmt"
	"time"
)

func main() {

	//Criando um canal
	canal := make(chan string)
	canal2 := make(chan string)

	go escrever("Olá mundo!", canal) //goroutine
	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		//Canal irá esperar receber o valor
		//Canal  está esperando receber um valor ao qual será salva na variável mensagem
		// o segundo parâmetro criado como 'open' determina para sabermos se o canal está em aberto ou não.
		mensagem, open := <-canal
		if !open {
			//break é o comando para encerrar a execução de um loop
			break
		}
		fmt.Println(mensagem)
	}

	go escrever2("Segunda função", canal2)
	fmt.Println("Depois da execução da segunda função")
	//Utilizando o comando range. Para cada mensagem recebida no canal
	for mensagem := range canal2 {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

//função que recebe uma string e recebe um canal
func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		//Enviar o dado pelo canal.
		canal <- texto
		time.Sleep(time.Second)
	}
	//Função para fechar o canal, para o mesmo não enviar e nem receber dados
	close(canal)

}

//função que recebe uma string e recebe um canal
func escrever2(texto string, canal2 chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		//Enviar o dado pelo canal.
		canal2 <- texto
		time.Sleep(time.Second)
	}
	//Função para fechar o canal, para o mesmo não enviar e nem receber dados
	close(canal2)
}
