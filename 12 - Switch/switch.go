package main

import (
	"fmt"
)

//Passando uma condição
//a cláusula default siginfica se nenhum case atender a necessidade, ele utiliza o default como retorno
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}

}

//Deixar a condição para ser avaliada em cada caso
//É mais útil quando não se quer avaliar a mesma varíavel
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

//Cláusula fallthrough
//Jogar o código para dentro da próxima condição. Não faz a avaliação da próxima condição
func diadaSemana3(numero int) string {
	var diaSem string

	switch {
	case numero == 1:
		diaSem = "Domingo"
		fallthrough
	case numero == 2:
		diaSem = "Segunda-Feira"
	case numero == 3:
		diaSem = "Terça-Feira"
	case numero == 4:
		diaSem = "Quarta-Feira"
	case numero == 5:
		diaSem = "Quinta-Feira"
	case numero == 6:
		diaSem = "Sexta-Feira"
	case numero == 7:
		diaSem = "Sábado"
	default:
		diaSem = "Número inválido"
	}
	return diaSem
}

func main() {
	fmt.Println("--------------Switch-------------")
	dia := diaDaSemana(10)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)

	dia3 := diadaSemana3(1)
	fmt.Println(dia3)
}
