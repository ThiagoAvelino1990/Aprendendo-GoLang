package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/*Cria um usuário no banco de dados*/
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	/*Leitura da requisição*/
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		/*Este return serve para parar o processamento*/
		return

	}

	/*Criando um usuário que está no pacote de modelos*/
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		/*Este return serve para parar o processamento*/
		return

	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		/*Este return serve para parar o processamento*/
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		/*Este return serve para parar o processamento*/
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	/*usuario que foi passado na requisição*/
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		/*Este return serve para parar o processamento*/
		return

	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

/*Buscar usuários no banco de dados*/
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)

	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

/*Buscar um usuário no banco de dados*/
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

/*Atualizar um usuário no banco de dados*/
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New(strings.ToUpper("você não tem permissão para realizar esta operação")))
		return
	}

	/*SE o usuário do login for diferente do usuário alterado, o processo será bloqeuado*/
	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New(strings.ToUpper("não é possível alterar um usuario que não seja o seu")))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

/*Deleta um usuário no banco de dados*/
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Excluindo usuário!"))
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New(strings.ToUpper("não é possível deletar um usuario diferente do seu")))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

/*FollowUser permite que um usuário siga outro*/
func FollowUser(w http.ResponseWriter, r *http.Request) {
	/*Extraur usuárioID do token*/
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/*Verifica se o usuário está seguindo ele mesmo*/
	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New(strings.ToUpper("Não é possível seguir você mesmo")))
		return
	}

	/*Conectar com o Banco de Dados*/
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	if erro = repositorio.Follow(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

/*UnfollowUser permite que um usuário deixa de seguir outro*/
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuariodID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuariodID {
		respostas.Erro(w, http.StatusForbidden, errors.New(strings.ToUpper("não é possível deixar de seguir você mesmo")))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	if erro = repositorio.Unfollow(usuariodID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInsufficientStorage, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

/*SearchFollowers tem como objetivo buscar os seguidores de um usuário*/
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	seguidores, erro := repositorio.SearchFollowers(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

/*SearchFollowing traz todos os usuários que um usuário está seguindo */
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeusuarios(db)
	usuarios, erro := repositorio.SearchFollowing(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

/*UpdatePassword serve para atualização de senha do usuário*/
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	/*Leitura do usuárioID que está no token*/
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	/*Leitura do usuarioID que está nos parâmetros*/
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/*Comparar os usuários(token x parametros)*/
	if usuarioIDNoToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New(strings.ToUpper("não é possível atualizar a senha de um usuário que não seja o seu")))
	}

	/*ler o corpo da requisição*/
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		/*Este return serve para parar o processamento*/
		return

	}

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/*Conexão com o banco de dados*/
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	/*Criando o repositorio*/
	repositorio := repositorios.NovoRepositorioDeusuarios(db)

	/*Validação da senha*/
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	/*Verificação das senhas*/
	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New(strings.ToUpper("senha atual divergente da senha salva no banco")))
		return
	}

	/*colocando o hash na nova senha*/
	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.UpdatePassword(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}
