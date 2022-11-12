package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*Esta função foi criada para apenas gerar uma única vez a secretkey
//função init roda antes da função main
func init() {
	//slice de byte de 64 posições
	chave := make([]byte, 64)

	//Utilizando o pacote rand para preencher o slice com valores aleatórios
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	//Convertendo o slice de byte em uma string base64
	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	//Imprimindo a string de base64. Com o resultado, basta colocar a variável SECRET_KEY no arquivo .env
	fmt.Println(stringBase64)

}
*/
func main() {
	config.Carregar()
	fmt.Println(config.Porta)
	fmt.Println(config.StringConexaoBanco)
	//fmt.Println(config.SecretKey)
	fmt.Println("Rodando API!")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
