package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end   endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	//Declaração variável explícita do tipo struct
	var u usuario
	//imprimindo valores zeros para o struct
	fmt.Println(u)

	//Populando a variável do tipo struct
	u.nome = "Doidin"
	u.idade = 31
	fmt.Println(u)

	//Criando variável devido a uma struct dentro de outra struct
	endereco := endereco{"Rua dos doidos", 40}

	//Declaração dos valores do struct com inferência de tipos
	//Necessário passar a variável criada para o segundo struct criado
	u2 := usuario{"Doidin2", 27, endereco}
	fmt.Println(u2)

	//Declaração dos valores com inferência de tipos sem passar todos os dados
	u3 := usuario{nome: "Doidin3"}
	fmt.Println(u3)

	//

}
