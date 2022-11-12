package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

/*ErroAPI representa a resposta de erro da API*/
type ErroAPI struct {
	Erro string `json:"erro"`
}

/*JSON retorna uma resposta em formato JSON para a requisição*/
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	/*tanto conteúdo dos dados que estamos mandando para a API é o application/json que informamos no controllers/usuarios quanto os dados que voltam*/
	w.Header().Set("Content-Type", "application/json")

	/*Para colocar no cabeçalho*/
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}

}

/*TratarStatsuCodeDeErro trata as requisições com statusCode 400 ou superior*/
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	/*Criado uma variável do tipo struct*/
	var erro ErroAPI
	/*transformando o erro em um struct*/
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)

}
