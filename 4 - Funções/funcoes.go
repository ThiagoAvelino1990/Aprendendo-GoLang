package main

import "fmt"

//Estrutura de uma função
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Função com mais de um retorno.
//Como os parâmetros são do mesmo tipo, basta informar o tipo do parâmetro no final
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	//Variável recebe valor de retorno da função
	soma := somar(10, 20)
	fmt.Println("Resultado da função 'somar' na variável 'soma': ", soma)

	//Declarar uma variável do tipo função e jogar uma função dentro da variável
	var f = func(txt string) string {
		fmt.Println("Texto da variável 'f' do tipo função: ", txt)
		return txt
	}

	//Variável que recebe o valor de uma variável do tipo função
	resultado := f("Texto da função 1")
	fmt.Println("Resultado da variável 'resultado', que recebe o valor da variável 'f' do tipo função: ", resultado)

	//Declaração de variável com funçõa.
	//Como a função exemplo contém dois retornos, deve-se declarar duas variáveis
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println("Variáveies declaradas que recebem os retorno da função 'calculosMatematicos' ", resultadoSoma, resultadoSubtracao)

	//Caso a função contenha mais de um retorno e você por algum motivo precise apenas um retorno é necessário utilizar o underline(_) na declaração das variáveis
	//Neste caso, a ordem dos fatores altera o produto, se quisesse apenas o outro resultado, o underline(_) ficaria no primeiro e não no segundo
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println("Imprimindo apenas um resultado da função 'calculosMatematicos' ", resultadoSoma2)

}
