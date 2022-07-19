package main

import "fmt"

//Função recursiva
//Funções que chamam elas mesmas, é uma função que dependesse de uma outra execução dela mesma
//Funções recursivas tem condição de parada, para não ficar em uma execução infinita. Oque chamamos de estouro de pilha
func fibonacci(posicao uint) uint {
	//condição de parada
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("------------FUNÇÕES RECURSIVAS------------")
	//Transformando a variável em UINT
	posicao := uint(7)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
