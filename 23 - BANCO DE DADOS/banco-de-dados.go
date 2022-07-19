package main

import (
	"database/sql"
	"fmt"
	"log"

	/*Forma para importar o pacote implícitamente*/
	/*Pacote necessário para a conexão com o MYSQL*/
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*Criação de uma variável para conexão com o banco de dados*/
	/*"OWNER:SENHA@/SERVICE_NAME?charset=utf8&parseTime=True&loc=Local"*/
	stringConexao := "GOLANG:GOLANG@/devbook?charset=utf8&parseTime=True&loc=Local"

	/*Função Open do pacote sql - recebe o banco e a string de conexao
	E retorna um ponteiro e um erro*/
	db, erro := sql.Open("mysql", stringConexao)

	/*Se o retorno da função sql.Open for diferente de null roda o IF*/
	if erro != nil {
		fmt.Println("Erro dentro do SQL Open")
		log.Fatal(erro)
	}

	defer db.Close()

	/*Reaproveitando a variável erro
	A variável erro recebe o retorno da função Ping, se for diferente de null roda o IF*/
	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("conexão está aberta!")

	/*Função Query retorna os dados do SELECT*/
	linhas, erro := db.Query("SELECT * FROM USUARIOS")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
