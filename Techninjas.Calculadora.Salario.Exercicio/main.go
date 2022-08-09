package main

import (
	"Packages/trabalhador"
	"fmt"
)

func main() {

	fmt.Println("Favor infromar o nome: ")
	fmt.Scanln(&trabalhador.Nome)

	fmt.Println("Favor informar o salario bruto: ")
	//Função Scanf utilizada para leitura dos dados informados pelo usuário
	fmt.Scanln(&trabalhador.SalarioBruto)

	//O comando %.2f indica a formatação de string para float com duas casas decimais após a vírgula
	//O comando \n indica pular linha, uma vez que está sento utilizando o fmt.Printf para impresão dos dados na tela
	fmt.Printf("Valor do desconto de IR: %.2f\n", trabalhador.CalcularIR(trabalhador.SalarioBruto))
	fmt.Printf("Valor do desconto de INSS: %.2f\n", trabalhador.CalcularINSS(trabalhador.SalarioBruto))

	fmt.Print("O Nome do trabalhador é: ", trabalhador.Nome, " Valor total do salario: ", trabalhador.CalculaSalarioLiquido(trabalhador.SalarioBruto))

}
