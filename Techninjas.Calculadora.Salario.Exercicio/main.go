package main

import (
	"Packages/trabalhador"
	"fmt"
)

func main() {
	salarioBruto := trabalhador.SalarioBruto
	//nomeTrabalhador := trabalhador.Nome

	fmt.Println("Favor informar o salario bruto: ")
	//Função Scanf utilizada para leitura dos dados informados pelo usuário
	fmt.Scanf("%f", &salarioBruto)

	//O comando %.2f indica a formatação de string para float com duas casas decimais após a vírgula
	//O comando \n indica pular linha, uma vez que está sento utilizando o fmt.Printf para impresão dos dados na tela
	fmt.Printf("Valor do desconto de IR: %.2f\n", trabalhador.CalcularIR(salarioBruto))
	fmt.Printf("Valor do desconto de INSS: %.2f\n", trabalhador.CalcularINSS(salarioBruto))

	fmt.Printf("Valor total do salario: %.2f\n", trabalhador.CalculaSalarioLiquido(salarioBruto))

}
