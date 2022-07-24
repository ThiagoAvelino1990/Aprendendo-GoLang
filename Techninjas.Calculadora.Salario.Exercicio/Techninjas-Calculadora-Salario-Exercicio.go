package main

import (
	"fmt"
)

//CalcularIR tem como objetivo calcular o descondo do IR do salário
func CalcularIR(salario float64) float64 {
	descIR := 0.0
	if salario >= 1903.99 && salario <= 2829.65 {
		descIR = (((salario * 7.5) / 100) - 142.80)
		return descIR
	}
	if salario >= 2826.66 && salario <= 3751.05 {
		descIR = (((salario * 15) / 100) - 354.80)
		return descIR
	}
	if salario >= 3751.06 && salario <= 4664.68 {
		descIR = (((salario * 22.5) / 100) - 636.13)
		return descIR
	}
	if salario >= 4664.68 {
		descIR = (((salario * 27.5) / 100) - 869.36)
		return descIR
	}

	return descIR
}

//CalcularINSS tem como objetivo calcular o desconto de INSS do salário
func CalcularINSS(salario float64) float64 {
	descINSS := 0.0
	if salario <= 1212.00 {
		descINSS = ((salario * 7.5) / 100)
		return descINSS
	}
	if salario >= 1212.01 && salario <= 2427.35 {
		descINSS = ((salario * 9) / 100)
		return descINSS
	}
	if salario >= 2427.36 && salario <= 3641.03 {
		descINSS = ((salario * 12) / 100)
		return descINSS
	}
	if salario >= 3641.04 && salario <= 7087.22 {
		descINSS = ((salario * 12) / 100)
		return descINSS
	}
	return descINSS
}

func main() {
	var salario float64

	fmt.Print("Informe o seu salário: ")
	//Função Scanf utilizada para leitura dos dados informados pelo usuário
	fmt.Scanf("%f", &salario)

	//Variável declarada por inferência de tipo
	desc := CalcularIR(salario) + CalcularINSS(salario)

	//O comando %.2f indica a formatação de string para float com duas casas decimais após a vírgula
	//O comando \n indica pular linha, uma vez que está sento utilizando o fmt.Printf para impresão dos dados na tela
	fmt.Printf("Valor do desconto de IR: %.2f\n", (CalcularIR(salario)))
	fmt.Printf("Valor do desconto de INSS: %.2f\n", CalcularINSS(salario))
	fmt.Printf("Valor total dos descontos %.2f\n", desc)

	fmt.Printf("Valor total do salario: %.2f\n", (salario - (CalcularIR(salario) + CalcularINSS(salario))))

}
