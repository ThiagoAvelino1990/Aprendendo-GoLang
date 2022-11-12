package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

/*Rota representa todas as rotas da aplicação web*/
type Rota struct {
	URI                string
	Metodo             string
	Funcao             http.HandlerFunc
	RequerAutenticacao bool
}

/*Configurar coloca todas as rotas dentro do router*/
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}

	/*Configuração para passar os comandos assets
	Desta forma não é necessário informar todo o caminho do arquivo css dentro do arquivo html, e nem também utilizando os dois ponto(..)*/
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
