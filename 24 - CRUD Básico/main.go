package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD - CREATE, RED, UPDATE, DELETE

	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE
	fmt.Println("CRUD -> CREATE(Criar dados); READ(Ler dados); UPDATE(Atualiza dados); DELETE(Exclui dados")

	//Irá conter todas as rotas para interações com o banco de dados
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Conectando...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
