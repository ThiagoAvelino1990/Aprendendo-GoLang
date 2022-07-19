package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("-----------Array e slice-----------")

	//Array é nada mais que uma lista de valores. todos os dados dentro do array devem ser do mesmo tipo
	var array1 [5]int
	//Como o array não tem valor, ele irá mostrar os 5 itens com o valor zero
	fmt.Println("Array do tipo 'int' com 5 posições: ", array1)

	//Imprimindo string nulo
	var array2 [5]string
	fmt.Println("Array do tipo 'string' com 5 posições: ", array2)

	//Informando um valor para uma posição do array
	array2[2] = "Posição 3"
	fmt.Println(array2)

	//Declarando array por inferência de tipo
	array3 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array3)

	//Fixando o tamanho do array de acordo com os valores passados
	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	//SLICE
	//Slice é uma estrutura que é baseada no array, porém ela é mais flexível principalemente no tamanho.
	//Ele não contém um valor fixo, ele aumenta de acordo com a necessidade.
	//Slice assim como um array, tem a limitação do tipo, ou seja, todos os valores devem ser do mesmo tipo
	//Não contém um tamanho fixo, o tamanho pode mudar de acordo com a necessidade
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//slice não é uma array, ele 'aponta' para um array. É como se fosse uma fatia de um array

	//reflect.TypeOf --> Função que devolve o tipo de uma variável
	//Slice não contém índices
	fmt.Println("Usando a função reflect.TypeOf() para o slice ", reflect.TypeOf(slice))
	//Array tem índices
	fmt.Println("Usando a função reflect.TypeOf() para o array ", reflect.TypeOf(array3))

	//Não tem como interagir entre array e slice.

	//Função append -> Adiciona um item novo no slice já criado
	slice = append(slice, 18)
	fmt.Println(slice)

	//Criando um pedaço do array e salvar dentro do slice
	//O primeiro índice é INCLUSIVO(inclui o primeiro índice), e o último índice é EXCLUSIVO(ignora o último índice)
	slice2 := array3[1:4]
	fmt.Println(slice2)

	//Alterando o valor de uma posição do array
	array3[2] = "Posição Alterada"
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	//Função 'make' aloca espaço na memória para alguma determinada coisa do programa
	fmt.Println("----------------Arrays Internos----------------")
	//Função 'make' cria array com tipo, tamanho e capacidade
	//Exemplo cria um array de 11 posições e retornar um slice de 10 posições(array interno)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	//Função len mostra o tamanho do slice
	fmt.Println(len(slice3)) //length
	//Função cap mostra a capacidade do slice
	fmt.Println(cap(slice3)) //capacidade

	//Adicionando um
	slice3 = append(slice3, 12)
	fmt.Println("Adicionando um para o length e cap ficarem iguais ", slice3)
	fmt.Println("length ", len(slice3)) //length
	fmt.Println("cap ", cap(slice3))    //capacidade

	//Quando se ultrapassa a capacidade de um slice, o GO dobra o tamanho
	slice3 = append(slice3, 13)
	fmt.Println("Estourando a capacidade ", slice3)
	fmt.Println("length ", len(slice3)) //length
	fmt.Println("cap ", cap(slice3))    //capacidade

	//Utilizando a função 'make' sem informar o cap. Desta forma, o array contém a length igual a cap
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println("length ", len(slice4)) //length
	fmt.Println("cap ", cap(slice4))    //capacidade

	//Adicionando para estourar a cap. Desta forma o GO dobra o tamanho
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println("length ", len(slice4)) //length
	fmt.Println("cap ", cap(slice4))    //capacidade
}
