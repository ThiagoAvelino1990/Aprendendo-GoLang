package main

import "fmt"

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
	fmt.Scanf("%f", &salario)

	desc := CalcularIR(salario) + CalcularINSS(salario)

	fmt.Printf("Valor do desconto de IR: %f\n", CalcularIR(salario))
	fmt.Printf("Valor do desconto de INSS: %f\n", CalcularINSS(salario))
	fmt.Printf("Valor total dos descontos %f\n", desc)

	fmt.Printf("Valor total do salario: %f\n", (salario - (CalcularIR(salario) + CalcularINSS(salario))))

}
