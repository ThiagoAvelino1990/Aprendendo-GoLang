package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o MYSQL
)

//Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	//string de conexão
	stringConexao := "GOLANG:GOLANG@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	//Utilizando a mesma variável erro declarada acima
	//Verifica se a conexão com o banco é efetivada
	if erro != nil {
		//Valor zero de ponteiros é sempre nil
		return nil, erro
	}
	//Utilizando a mesma variável erro declarada acima
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	//Se a conexão deu certo, retornamos com a conexão e o erro vazio
	return db, nil

}
