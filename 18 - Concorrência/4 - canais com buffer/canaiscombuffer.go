//Criar um canal com buffer é específicar a capacidade do canal;
//O Canal com buffer só bloquea quando atingir a capacidade máxima dele;
package main

import "fmt"

func main() {
	//Criando um canal com buffer
	//Neste caso, o buffer declarado é de '2'. E está se passando dois valores, se por algum motivo informar mais um valor, o programa não terá erro de compilação,
	// mas sim de execução
	canal := make(chan string, 2)
	//enviando valor para o canal
	canal <- "Canal com buffer"
	canal <- "Capacidade de 2"
	//canal <- "Erro de deadlock"

	//Recebendo o valor do canal
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
