package main

import "fmt"

func main() {
	fmt.Println("-----------MAP-----------")

	//Dentro dos colchetes é o tipo das chaves e fora dos colchetes é o tipo dos valores
	//Sempre que declarar uma chave, deve-se colocar vírgula(,) se não o GO não compila o código
	usuario := map[string]string{
		"Nome":      "Pedro",
		"Sobrenome": "Doidin",
	}

	fmt.Println(usuario)

	//Acessando uma chave do map
	fmt.Println(usuario["Nome"])

	//Modo para aninhar um map
	usuario2 := map[string]map[string]string{
		"Nome": {
			"FirstName": "Pedro",
			"LastName":  "Cabeção",
		},
		"Curso": {
			"NomeCurso": "Engenharia",
			"Campus":    "Party",
		},
	}

	fmt.Println(usuario2)

	//Apagando uma chave do map
	delete(usuario2, "Nome")
	fmt.Println(usuario2)

	//Adicionar uma chave do map
	usuario2["Nome"] = map[string]string{
		"FirstName": "João",
		"LastName":  "José",
	}

	fmt.Println(usuario2)
}
