package main

import "fmt"

func alunoEstaAprovado(n1, n2 uint) bool {
	defer recuperarExecução()

	if ((n1 + n2) / 2) > 6 {
		return true
	} else if ((n1 + n2) / 2) < 6 {
		return false
	}
	//Declarando a função PANIC. Sempre deve passar algum parâmetro para que a mesma seja executada
	//A função PANIC encerra imediatamente o programa, porém ela executa todas as funções com DEFER antes de fechar o programa
	panic("A média é exatamente igual a 6")

}

func recuperarExecução() {
	//Utilizando a função RECOVER, para recuperar a execução do programa
	//A declaração da função RECOVER deve ser feita com um IF INIT, aonde a variável recebe o falor da função RECOVER, e a mesma deve ser diferente de NULO
	if r := recover(); r != nil {
		fmt.Println("A execução foi recuperada com sucesso")
	}
}

func main() {
	fmt.Println("------------PANIC e RECOVER------------")
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução realizada com sucesso")

	//Não há problema usar a função recover em uma função que não irá ter uma cláusula PANIC
}
