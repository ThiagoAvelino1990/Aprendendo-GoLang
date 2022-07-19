package main

import "fmt"

//Criando uma função que recebe como parâmetro uma interface genérica
func generica(i interface{}) {
	fmt.Println(i)
}

//Declarando uma interface genérica
type intGenerica interface {
}

//Crianção função para receber a interface genérica
func generica2(i intGenerica) {
	fmt.Println(i)
}

func main() {

	//Qualquer informação passada para a interface, ele irá executar sem erro
	generica("Doido do jipe")
	generica(30)
	generica(false)

	//Declaração de variável explícita
	var b string
	b = "Testando"
	generica2(b)

	//interface generico utilizando em um MAP
	//o tipo do map pode ser qualquer tipo de interface, bem como o retorno
	mapa := map[interface{}]interface{}{
		1:        10,
		"String": "String",
		true:     float64(12),
	}

	fmt.Println(mapa)
}
