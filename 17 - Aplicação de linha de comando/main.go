package main

//Importação dos pacotes
import (
	"fmt"
	//Pacote importado: nome_do_modulo/pacote_criado
	"linhadecomando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Ponto de Partida")

	//Variável aplicacao recebe a função Gerar do pacote app
	aplicacao := app.Gerar()
	//os.Args é um parâmetro padrão para que os comandos sejam reconhecidos
	//IF init criado para tratamento do erro
	if erro := aplicacao.Run(os.Args); erro != nil {
		//Pacote log faz algo parecido com o pacote fmt
		//essa declaração log.Fatal, se desse algum erro faria com que ele parasse a execução do programa
		log.Fatal(erro)
	}

}
