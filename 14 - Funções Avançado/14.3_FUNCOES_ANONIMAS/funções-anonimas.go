package main

import "fmt"

func main() {
	//Funções anonimas
	//Basicamente é uma função que não tem nome. Não delcara a função explícitamente
	//Para executar a função, deve-se colocar parênteses() no final da função
	func() {
		fmt.Println("Primeiro exemplo de função anônima")
	}()

	//Informando o parâmetro para a função fazer o processo dela
	func(texto string) {
		fmt.Println(texto)
	}("Segundo exemplo de função anônima") //Os parâmetros de uma função anônima, devem ser passados por aqui

	//Colocando retorno para a função anônima
	//Deve-se declarar uma variável para armazenar o retorno da função
	retorno := func(texto string) string {
		return fmt.Sprintf("Terceiro exemplo de função anônima -> %s", texto)
	}("Valor do parâmetro")

	fmt.Println(retorno)

	//Treinando a crianção de uma função anônima
	fmt.Println(func(n1, n2 int) (soma int, sub int) {
		soma = n1 + n2
		sub = n1 - n2
		return

	}(2, 5))

}
