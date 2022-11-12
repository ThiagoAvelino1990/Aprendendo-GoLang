package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	/*StringConexaoBanco - String de conexão com o banco MYSQL*/
	StringConexaoBanco = ""
	/*Porta aonde a  API vai estar rodando*/
	Porta = 0

	//SecretKey chave para assinatura do token do tipo slice de byte
	SecretKey []byte
)

/*Carregar variáveis de ambiente*/
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	/*Buscando a variável do arquivo .env*/
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	/*Passando os dados de arquivo de conexão*/
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	/*Buscando o valor da chave SECRET_KEY no arquivo .env*/
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
