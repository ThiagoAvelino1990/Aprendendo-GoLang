package main

import (
	"log"
	"net/http"
)

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func golang(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Aprendendo GOLANG"))
}

func main() {
	/*
		HTTP é um protocolo de comunicação = Base de comunicação WEB
		CLIENTE (FAZ REQUISIÇÃO - REQUEST) - SERVIDOR (PROCESSA A REQUISIÇÃO E ENVIA A RESPOSTA - RESPONSE)
		ROTAS É MANEIRA DE IDENTIFICAR O TIPO DE MENSAGEM ESTÁ SENDO ENVIADA E IDENTIFICAR QUE TIPO DE PROCESSAMENTO DEVE SER FEITO EM CIMA DESTA MENSAGEM
		      -COMPONENTE URI - IDENTIFICADOR DO RECURSO(Uma forma de informar ao servidor sobre oque está sendo passado)
			  - MÉTODO - Usado para informar oque deseja ser feito para o recurso identificado
			  	*GET(BUSCAR DADOS DE UM RECURSO), POST(CADASTRAR DADOS), PUT(ATUALIZAÇÕES DE DADOS), DELETE(DELETAR DADOS DA BASE)

	*/
	/*
		Declaração do MÉTODO HTTP
		Função HandleFunc recebe dois parâmetros:
			1 - URI da rota
			2 - É uma função que irá tratar a requisição(request).
				Dentro desta função devemos passar dois tipos de parâmetros pois, ela deve ser reconhecida pela função Handlefunc:
					1 - Do tipo http.ResponseWritter(Response de volta para o cliente que fez a requisição)
					2 - Um ponteiro do tipo *http.Request(Requisição que está sendo enviada, por "padrão" o mesmo é referênciado como r)
						Como é HTTP não podemos usar a função fmt.Println() para escrever uma mensagem, devemos utilzar o Write([]byte) que escreve uma mensagem e a mesma deve estar obrigatóriamente em um slice de byte

		Neste caso em específico, para a utilização basta rodar o arquivo .go e no navegador digitar locallhost -> pois o servidor é a própria máquina
																									 :5000 		->porta definina na função log.Fatal(http.ListenerAndServer())
																									 /home		-> URI definina na função HandleFunc/home
																									localhost:5000/home

	*/
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	//localhost:5000/usuarios
	http.HandleFunc("/usuarios", usuarios)

	//localhost:5000/golang
	http.HandleFunc("/golang", golang)

	/***
	Comando para criação de um servidor
	Primeiro parâmetro é uma porta.	Porta essa responsável por REQUEST e RESPONSE
	***/
	log.Fatal(http.ListenAndServe(":5000", nil))

}
