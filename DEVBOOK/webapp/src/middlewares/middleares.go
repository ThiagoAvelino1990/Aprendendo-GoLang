package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

/*Logger escreve informações da requisição no terminal*/
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s, %s, %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}

}

/*Autenticar Verifica a existência de cookies
Está função  não valida se os dados estão corretos.
Ela apenas verifica se os dados de autenticação existem no browser do usuário, no cookies
A autenticação é feita pela API*/
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*Valida se os dados da autenticação estão dentro do cookie*/
		if _, erro := cookies.Ler(r); erro != nil {
			/*caso ocorra erro, será redirecionado para a rota de login*/
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		proximaFuncao(w, r)
	}
}
