package main

import (
	"fmt"
)

func main() {

	//OPREADORES ARITIMÉTICOS
	//(+);(-);(/);(*);(%)
	soma := 10 + 5
	subtracao := 10 - 5
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 5
	fmt.Println("Soma: ", soma, "Subtracao: ", subtracao, "Divisao: ", divisao, "Multiplicação: ", multiplicacao, "Resto da divisão(MOD): ", restoDaDivisao)
	fmt.Println("--------------------")

	//OPERADORES DE ATRIBUIÇÃO
	// = -> Com o igual(=) é necessário declarar o tipo da variável
	//:= -> Inferência de tipos
	var string1 string = "String 1"
	string2 := "String 2"
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println("--------------------")

	//OPERADORES RELACIONAIS
	//Operadores relacionais sempre retornam true ou falso(resultado booleando)
	fmt.Println(1 > 0)  //Maior >
	fmt.Println(1 < 0)  //Menor <
	fmt.Println(1 >= 0) //Maior ou igual >=
	fmt.Println(1 == 0) //Igual(Comparação) ==
	fmt.Println(1 <= 0) //Menor ou igual <=
	fmt.Println(1 != 0) //Diferente !=
	fmt.Println("--------------------")

	//OPERADORES LÓGICOS
	//&& (e) - Se conter um valor falso e um valor verdadeiro, o retorno é falso
	//|| (ou)- Se conter um valor falso e um valor verdadeiro, o retorno é verdadeiro
	//! (Not)- Este operador nega o valor da variável booleanda. Se for verdadeiro, o retorno é falso e vice-versa

	verdadeiro := true
	falso := false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println("--------------------")

	//OPERADORES UNÁRIOS
	//Operadores que só interagem com uma variável por vez. Estes operadores devem vir sempre após a variável
	numero := 10
	numero++ //++ Incrementar + 1 (numero = numero + 1)
	fmt.Println(numero)
	numero += 10 //+= Incrementar + (Valor desejado) (numero = numero + 10)
	fmt.Println(numero)
	numero-- //-- Decrementar - 1 (numero = numero - 1)
	fmt.Println(numero)
	numero -= 10 //-= Decrementar - (Valor desejado) (numero = numero - 10)
	fmt.Println(numero)
	numero *= 2 //*= Incrementar * (Valor desejado) (numero = numero * 2)
	fmt.Println(numero)
	numero /= 10 // /= Divisão / (Valor desejado) (numero = numero / 10)
	fmt.Println(numero)
	numero %= 10 //%= MOD (Valor desejado) (numero = numero % 10)
	fmt.Println(numero)
	fmt.Println("--------------------")

	//OPERADORES TERNÁRIOS
	//Operadores que muitas linguagens usam para colocar o valor de uma variável de acordo com uma condição
	//Porém, não existe este tipo de operadores na linguagem GO. Para realizar este tipo de operação, deve-se utilizar o IF e o ELSE
	var texto string
	if numero > 5 {
		texto = "Mario que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
