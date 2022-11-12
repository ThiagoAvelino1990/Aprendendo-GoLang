package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/*A função init serve para cerregar os templates antes da execução da função main*/
func init() {
	utils.CarregarTemplates()
}

/*Função INIT apenas para gerar a secretKey
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)

}
*/

func main() {
	/*Carregando variáveis de ambiente*/
	config.Carregar()
	cookies.Configurar()

	r := router.Gerar()

	fmt.Printf("Rodando WebAPP! Na porta %d\n", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
