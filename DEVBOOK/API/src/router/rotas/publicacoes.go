package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	/*Rota para Criar Publicações*/
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	/*Rota para buscar Publicações. Não irá buscar todas as publicações salvas no banco. Quando um usuário chamar essa rota, ele irá receber todos as publicações de quem ele segue, e as suas próprias publicações*/
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	/*Rota para buscar uma única publicação*/
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	/*Rota para Atualizar Publicação*/
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	/*Rota para deletar uma Publicação*/
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	/*Rota para buscar publicacoes de um usuario*/
	{
		URI:                "/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},
	/*Rota para like uma publicacao*/
	{
		URI:                "/publicacoes/{publicacaoId}/like",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	/*Rota para dislike uma publicacao*/
	{
		URI:                "/publicacoes/{publicacaoId}/dislike",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
