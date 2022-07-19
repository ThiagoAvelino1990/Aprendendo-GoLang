package main

//Importação dos pacotes
import (
	"fmt"
	//Pacote criado com as funções Escrever() e escrever2(). Para importar deve informar o nome do módulo criado e o nome do pacote
	"modulo/auxiliar"

	//Pacote externo para checkar email(GitHub)
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	//Para executar a função abaixo, deve-se utilizar o nome do pacote
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
