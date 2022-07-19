package main

import "fmt"

//Função normal precisa de um return pois a função apenas copia o valor da variável e faz a operação. Aaltera o valor da variável somente na execução da função.
func inverterSinal(n int) int {
	return n * -1
}

//função com ponteiro não precisa de um retorno pois ela irá alterar o valor de acordo com o local na memória para a variável
func inverterSinalComPonteiro(n *int) {
	*n = *n * -1
}
func main() {
	fmt.Println("Função com ponteiro")

	numero := 20
	fmt.Println("Utilizando a função", inverterSinal(numero))
	fmt.Println("Após a execução da função, o valor da variável não muda", numero)

	novoNumero := 40
	fmt.Println("Valor da variável antes da execução da função", novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Valor da variável após a execução da função", novoNumero)

}
