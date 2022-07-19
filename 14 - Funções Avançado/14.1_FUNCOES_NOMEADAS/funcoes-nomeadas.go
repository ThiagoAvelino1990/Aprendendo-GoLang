package main

import "fmt"

//Funções com retorno nomeado
//Declarando as variáveis no retorno da função, não é necessário informá-las na cláusula return
func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	fmt.Println("----------Funções Nomeadas----------")
	fmt.Println(calculosMatematicos(10, 5))

}
