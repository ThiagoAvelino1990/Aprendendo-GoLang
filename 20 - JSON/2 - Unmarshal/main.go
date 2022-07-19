package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // chave representada dentro de um JSON
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//Declaração de uma variável para utilizar na função Unmarshal.
	//a Variável é do tipo string porém está em um formato legível
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	//Criação de variável do tipo struct em branco
	c := cachorro{}
	fmt.Println("Struct Vazio ", c)

	/*
		A função json.Unmarshal recebe dois parâmetros: 1º slice de byte. 2º endereço de memória
		Neste caso,para transformar a variável em um slice de byte é necessário a seguinte sintaxe []byte("nome_var")
		O segundo parâmetro deve-se colocar o e comercial(&) para referenciar a endereço de memória
		Esta função retorna somente um valor, que seria um erro
	*/
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	} else {
		//Impressão no formato struct e com os dados preenchidos
		fmt.Println(c)
	}

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`

	//Declaração de variável do tipo MAP
	c2 := make(map[string]string)
	fmt.Println("Map vazio ", c2)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	} else {
		//Impressão no formato struct e com os dados preenchidos
		fmt.Println(c2)
	}

}
