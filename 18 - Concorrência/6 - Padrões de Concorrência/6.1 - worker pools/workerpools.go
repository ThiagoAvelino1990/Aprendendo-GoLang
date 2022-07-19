//Padrões de concorrencia: São maneiras que você pode aplicar concorrência em seu código, seguinte algumas práticas ou metodologias que podem fazer
// que você tenha um resultado muito bom em relação as funções
//worker pools: filas de tarefas para serem executadas. E vários workers, processos pegando itens dessa fila de maneira independente
package main

import "fmt"

//Função recursiva
//Funções que chamam elas mesmas, é uma função que depende de uma outra execução dela mesma
//Funções recursivas tem condição de parada, para não ficar em uma execução infinita. Oque chamamos de estouro de pilha
func fibonacci(posicao int) int {
	//condição de parada
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//Sintaxe para o canal de tarefas só recebe dados
//Sintaxe para o canal de resultados só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	//intera
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	fmt.Println("------------FUNÇÕES RECURSIVAS COM WORKER POOLS------------")
	//Canel do tipo int com buffer de 45
	tarefas := make(chan int, 45)    //Canal que irá conter a sequencia de números
	resultados := make(chan int, 45) //Canal que irá armazenar os resultados a medida que os resultados forem gerados

	//Executando funções worker
	//Processos que estarão realizando o mesmo processo para ganho de performance
	go worker(tarefas, resultados) //goroutine
	go worker(tarefas, resultados) //goroutine
	go worker(tarefas, resultados) //goroutine
	go worker(tarefas, resultados) //goroutine

	//for para preencher o canal de tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	//fechar o canal de tarefas
	close(tarefas)

	//For para impressão dos resultados na tela
	for i := 0; i < 45; i++ {
		//Variável resultado recebe o valor que será enviado para o canal resultados
		resultado := <-resultados
		//imprimindo o valor resultado
		fmt.Println(resultado)
	}

}
