package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

/*FazerRequisicaoComAutenticacao é utilizado para colocar o token na requisição*/
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	/*request apenas para ler o cookie. Aqui, a aplicação está criando a requisição*/
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	/*realizado a leitura do cookie ignorando o erro
	ignorando pois, já foi passado pelo middleware*/
	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	/*Realizando a requisição*/
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil

}
