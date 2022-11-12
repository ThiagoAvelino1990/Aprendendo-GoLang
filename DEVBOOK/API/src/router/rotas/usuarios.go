package rotas

import (
	"api/src/controllers"
	"net/http"
)

/*struct*/
var rotasUsuarios = []Rota{
	/*Cria Usuário*/
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	/*Busca usuários*/
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	/*Busca um usuário por ID*/
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	/*Atualiza um usuário*/
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	/*Deleta um usuário*/
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	/*Rota para um usuário seguir outro usuário*/
	{
		URI:                "/usuarios/{usuarioId}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FollowUser,
		RequerAutenticacao: true,
	},
	/*Rota para um usuário deixar de seguir outro usuário*/
	{
		URI:                "/usuarios/{usuarioId}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UnfollowUser,
		RequerAutenticacao: true,
	},
	/*Rota para buscar os seguidores*/
	{
		URI:                "/usuarios/{usuarioId}/followers",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchFollowers,
		RequerAutenticacao: true,
	},
	/*Rota para buscar os usuarários que um determinado usuário está seguindo*/
	{
		URI:                "/usuarios/{usuarioId}/following",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchFollowing,
		RequerAutenticacao: true,
	},
	/*Rota para atualização de senha*/
	{
		URI:                "/usuarios/{usuarioId}/updatepassword",
		Metodo:             http.MethodPost, /*Método normalmente utilizado para atualização de senha pois o PUT ao receber duas requisições passando os mesmos dados a segunda alteração não terá o mesmo efeito*/
		Funcao:             controllers.UpdatePassword,
		RequerAutenticacao: true,
	},
}
