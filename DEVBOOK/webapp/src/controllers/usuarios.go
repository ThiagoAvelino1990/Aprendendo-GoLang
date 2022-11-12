package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

/*CriarUsuario chama a API para cadastrar um usuário no banco e dados*/
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	/*Coletando os dados do formulário web*/
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	/*Requisição para a API*/
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	/*Bloco para tratamento de mensagem de erro*/
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	/*o nil é necessário pois não é preciso passar os dados do usuário*/
	respostas.JSON(w, response.StatusCode, nil)

}
