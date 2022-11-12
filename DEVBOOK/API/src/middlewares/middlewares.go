package middlewares

/*********************************************************************
Camada que fica entre requisição e respostas.
Muita utilizado nas rotas, e para não ficar criado uma para cada rota,
normalmente se cria um pacote middleware.
Serve para aninhamento de funções
*********************************************************************/

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

/*Logger escreve informações da requisição no terminal*/
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

/*Autentucar verifica se o usuário fazendo a requisição está autenticado*/
/*Tipo HandlerFunc é igual a (http.ResponseWriter, *http.Request)*/
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r) /*Executa a função que vem pelo parâmetro*/
	}
}
