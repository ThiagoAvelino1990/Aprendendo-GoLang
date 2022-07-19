package main

import "fmt"

//Funçãos são tipos em GO, então podemos retornar elas, passar como parâmetros e tudo mais
func clousure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	fmt.Println("----------Função Closure ----------")

	texto := "Dentro da função main"
	fmt.Println(texto)

	//Variável funcaoNova do tipo FUNÇÃO
	funcaoNova := clousure()
	//Chamando a variável funcaoNova
	funcaoNova()

}
