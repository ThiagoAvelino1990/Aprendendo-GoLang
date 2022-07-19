package main

import "fmt"

//Recebe n parâmetros, ou seja, não é necessário especificar a quantidade de parâmetros que essa função pode receber.
//Por criar um SLICE, é possível interar a função
func soma(n1 ...int) int {
	totalSoma := 0
	//Interando a função
	for _, numeros := range n1 {
		totalSoma += numeros
	}
	return totalSoma

}

//Criando uma função com um valor fixo, e um valor variático
//Não se pode ter mais de um parâmetro variático por função. E também, este tipo de parâmetro deve ser o último
func escrever(texto string, n1 ...int) {
	for _, numero := range n1 {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções variáticas")

	//Passando vários valores para a função
	valorTotal := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(valorTotal)

	//Chamando a função com valor fixo e com valores variáticos
	escrever("Olá Mundo", 3, 422, 30)
}
