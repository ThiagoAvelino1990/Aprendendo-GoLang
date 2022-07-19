/*
JSON - Método utilizado para trafegar dados.
	 - Formado para armazenar dados ou transportar dados de um lado para o outro
	 - Muito utilizado para transportar dados de um servidor para uma página web e vice-versa
	 - Bem parecedio com struct no GO
*/
package main

import (
	"bytes"
	"encoding/json" //pacote para a utilização de JSON
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // chave representada dentro de um JSON
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//variável c do tipo struct cachorro
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	//json.Marshal - utilizado para converter um struct/ map para arquivo JSON
	//Transformando struct em JSON
	//Variável cachorroEmJSON recebendo um slice de byte
	//A função Marshal retorna um slice de byte e um erro
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	} else {
		//Imprimi um slice de byte
		fmt.Println(cachorroEmJSON)
		//pacote bytes com a função NewBuffer converte o slice em JSON
		fmt.Println(bytes.NewBuffer(cachorroEmJSON))
	}

	//Utilizando MAP para JSON
	c2 := map[string]string{
		"nome": "Tobby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	} else {
		fmt.Println(cachorro2EmJSON)

		fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
	}

}
