package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaPaginaPrincipal = Rota{
	/*Rota para carregar a página principal*/
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
