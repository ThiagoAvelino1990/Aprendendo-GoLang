package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func mediaDoAluno(n1, n2 float64) bool {
	//Com o defer abaixo, esta mensagem será executada antes do return da função
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	if ((n1 + n2) / 2) >= 6 {

		return true
	}
	return false
}

func main() {
	fmt.Println("---------------DEFER---------------")
	//Claúsula DEFER
	//Serve para adiar a retorno de alguma função, ou adiar a execução de um bloco de comandos.
	//É muito mais usando quando se está utilizando uma conexão com o Banco de Dados

	defer funcao1()

	funcao2()

	fmt.Println(mediaDoAluno(4, 6))

}
