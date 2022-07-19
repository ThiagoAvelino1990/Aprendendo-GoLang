package main

import "fmt"

func main() {
	fmt.Println("------------Estruturas de Controle------------")

	numero := 10

	//Sintaxe IF
	if numero > 5 {
		fmt.Println("numero é maior que 5")
	}

	//Sintaxe IF com ELSE
	if numero > 15 {
		fmt.Println("numero maior que 15")
	} else {
		fmt.Println("numero menor que 15")
	}

	//Sintaxe ELSE IF
	if numero > 20 {
		fmt.Println("numero maior que 20")
	} else if numero > 15 {
		fmt.Println("numero maior que 10")
	} else {
		fmt.Println("numero é igual a:", numero)
	}

	//Declarando uma variável dentro do IF com inferência de tipos
	//Variáveis declaradas dentro do IF não podem ser utilizadas fora do escopo. Após o término do IF, a variável "morre"
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("outronumero maior que 0. Valor de outroNumero:", outroNumero)
	} else {
		fmt.Println("outroNumero menor que 0")
	}

}
