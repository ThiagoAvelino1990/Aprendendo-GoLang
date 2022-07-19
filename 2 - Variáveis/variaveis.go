package main

import "fmt"

func main() {
	//Declaração de variável do tipo explícito, aonde informa o tipo e o valor da variável
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	//Declaração de variável do tipo implícito. Desta forma o GO determina qual o tipo da variável através do valor
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	//Declarando mais de uma variável do tipo explícito
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	//Declarando o valor de mais de uma variável do tipo implícito
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//Constante não se pode mudar o valor da constante. Ela tem as mesmas regras de declaração das variáveis
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//Inverter valor das variáveis(Em muitas linguagens seria necessário conter uma variável auxíliar)
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println("Trocando Valores: ", variavel5, variavel6)

}
