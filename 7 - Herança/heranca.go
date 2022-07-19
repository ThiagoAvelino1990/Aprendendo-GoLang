package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	//Declaração do struct 'pessoa' sem o tipo
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Joao", "Pedro", 20, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	fmt.Println("'Pseudo-Hereança'. Imprimindo a altura do struct 'pessoa' através da struct 'estudante'")
	fmt.Println(e1.altura)

}
