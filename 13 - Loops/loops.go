package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------LOOPS----------")

	//LOOPS
	//O GO utiliza somente o FOR como estrutura de repetição. não existe WHILE, DOWHILE, entre outros.

	//1ª maneira de usar o FOR, sintaxe com WHILE
	//Equanto uma condição for verdadeira, eu vou repetir o bloco de código
	fmt.Println("----------1ª maneira de usar o FOR----------")
	var i int = 1
	for i < 10 {
		i++
		//pode ser usada para formatar datas, pausar o código por alguns segundos entre outras
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	//2ª maneira de usar o FOR, sintaxe parecida com o IF INIT
	//Variáveis declaradas dentro do FOR não podem ser usadas fora do loop
	fmt.Println("----------2ª maneira de usar o FOR----------")
	for j := 1; j < 10; j++ {
		fmt.Println("Infrementando j", j)
		time.Sleep(time.Second)

	}

	//Cláusla Range, utilizar quando irá iterar em algo.(Array, String, Map e etc)
	//Exemplo: Declarar um array
	array := [3]string{"Posição A", "Posição B", "Posição C"}
	fmt.Println("----------Utilizando a cláusula range com array para iterar----------")
	//O primeiro por padrão é sempre o índice, e o segundo o valor
	for i_array, v_array := range array {
		fmt.Println(i_array, v_array)
		time.Sleep(time.Second)
	}

	//Iterando com slice
	slice := []string{"Posição", "Posição", "Posição", "Posição", "Posição", "Posição"}
	fmt.Println("----------Iterando com slice----------")
	for _, v_slice := range slice {
		fmt.Println(v_slice)
		time.Sleep(time.Second)
	}

	//Não usando o índice
	fmt.Println("----------Usando range com array sem informar o índice----------")
	for _, v_array := range array {
		fmt.Println(v_array)
		time.Sleep(time.Second)
	}

	//Iterar por STRING
	//Usando o range, ele irá trazer por padrão os valores da tabela ASCI de cada lertra
	fmt.Println("----------Iterar uma string----------")
	for i_string, v_string := range "PALAVRA" {
		//Usando a função(string) para trazer o valor correto
		fmt.Println(i_string, v_string, string(v_string))
		time.Sleep(time.Second)

	}

	//Iterar usando MAP
	fmt.Println("----------Iterar usando MAP----------")
	usuario := map[string]string{
		"Nome":      "Thiago",
		"Sobrenome": "Avelino",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
		time.Sleep(time.Second)
	}
	//Não se pode iterar no struct

	/*-----Por curiosidade, para criar um for infinito-----*/
	/*
		for {
			"condição"
		}

		for true {
			"condição"
		}
	*/
}
