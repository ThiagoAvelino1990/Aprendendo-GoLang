package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

/*JSON retorna uma resposta em JSON para a requisição*/
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	/*Repostas em JSON ao realizar a busca*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

/*Retorna um erro em JSON*/
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		/*Para preencher os dados do struct*/
		Erro: erro.Error(), /*Método Error dentro do tipo error, irá basicamente enviar um retorno de erro*/
	})

}
