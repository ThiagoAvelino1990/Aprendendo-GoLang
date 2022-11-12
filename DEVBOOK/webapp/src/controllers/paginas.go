package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"webapp/src/config"
	"webapp/src/respostas"
	"webapp/src/utils"
)

/*CarregarTelaDeLogin vai renderizar a tela de login*/
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

/*CarregarPaginaDeCadastroDeUsuarios irá carregar a página de cadastro de usuário*/
func CarregarPaginaDeCadastroDeUsuarios(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

/*CarregarPaginaPrincipal tem como objetivo carregar a página principal com as publicações*/
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%spublicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(r.Response.Body).Decode(&publicacoes); erro != nil {
		respotas.JSON(w, http.StatusUnprocessableEntinity, respotas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "home.html", publicacoes)
}
