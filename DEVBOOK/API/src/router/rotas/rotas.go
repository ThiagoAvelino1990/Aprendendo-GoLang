package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

/*Struct que representa todas as rotas da API*/
type Rota struct {
	URI                string
	Metodo             string
	Funcao             http.HandlerFunc /*func(http.ResponseWriter, *http.Request)*/
	RequerAutenticacao bool
}

/*Configurar coloca todas as rotas dentro do router*/
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...) /*criando rota para publicação. O formato é diferente, com os três pontos no final, pois é algo que sempre será adicionado*/

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
