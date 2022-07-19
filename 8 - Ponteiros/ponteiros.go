package main

import "fmt"

func main() {
	fmt.Println("------Ponteiros-----")

	var var1 int = 10
	var var2 = var1
	//A partir desta linhas, as variáveis não tem ligação nenhuma com a outra
	fmt.Println("Var1: ", var1, "Var2: ", var2)

	//Alterando apenas o valor da 'var1'. Atribuindo um valor a uma variável, este tipo é uma atribuição de valor por cópia.
	var1++
	fmt.Println("Var1: ", var1, "Var2: ", var2)

	//Ponteiro é uma referência de memória. Quando se cria um ponteiro, você atribui o endereço de memória aonde a variável está salva, e não o valor da variável.

	var var3 int
	//Para a declaração tem a cláusua 'var' e para identificar que é um ponteiro, deve-se utilizar asterísco(*)
	var ponteiro *int

	fmt.Println("var3 sem atribuição:", var3, "ponteiro sem atribuição:", ponteiro)

	var3 = 100
	//Para a atribuição o endereço de memória da variável deve-se usar 'e comercial'(&)
	ponteiro = &var3

	fmt.Println("var3:", var3, "ponteiro:", ponteiro)

	//Desreferenciação de um ponteiro também é usado no asterísco(*)
	fmt.Println("var3:", var3, "Colocando asterísco(*) para trazer o valor destinado no endereço de memória:", *ponteiro)

	//Altereando o valor da variável var3 o ponteiro não muda pois, continua referenciando o mesmo endereço de memória
	var3 = 150
	fmt.Println("var3:", var3, "ponteiro:", ponteiro)

}
