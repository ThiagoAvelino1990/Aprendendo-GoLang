package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

/*Não será exportado*/
type usuarios struct {
	db *sql.DB
}

/*Função cria um repositório de usuários*/
func NovoRepositorioDeusuarios(db *sql.DB) *usuarios {
	return &usuarios{db}

}

/*Criação do método*/
/*Criar insere um usuário no banco de dados*/
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO USUARIOS(NOME, NICK, EMAIL, SENHA) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

/*Assinatura do método*/ /*Parâmetros que recebe*/ /*Retorno do método*/
/*Buscar traz todos os usuários que atendem um filtro de nome ou nick*/
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) /* %nomeOuNick% */

	linhas, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM USUARIOS WHERE NOME LIKE ? OR NICK LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	/*Iterar sobre as linhas que retornaram do banco*/
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		/*repassando o slice de usuarios e a variável usuario*/
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

/*BuscarPorID traz um usuário do banco de dados*/
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM USUARIOS WHERE ID = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

/*Atualizar altera os dados de um usuário no banco de dados*/
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE USUARIOS SET NOME = ?, NICK = ?, EMAIL =? WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil

}

/*Deletar exclui as informações de um usuário do banco de dados*/
func (repositorio usuarios) Deletar(ID uint64) error {
	statemant, erro := repositorio.db.Prepare(
		"DELETE FROM USUARIOS WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statemant.Close()

	if _, erro = statemant.Exec(ID); erro != nil {
		return erro
	}

	return nil

}

/*BuscarPorEmail busca um usuario por email e retorna o seu ID e senha*/
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(`
	SELECT ID, SENHA 
	  FROM USUARIOS
	 WHERE EMAIL = ?
	`, email)
	/*Verificar se existe algum erro*/
	if erro != nil {
		/*Se existir retorna um usuário vazio, e o erro*/
		return modelos.Usuario{}, erro
	}
	/*fechar a linha*/
	defer linha.Close()

	var usuario modelos.Usuario
	/*iterar a variável linha*/
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

/*Follow permite que um usuário siga outro*/
func (repositorio usuarios) Follow(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO SEGUIDORES (USUARIO_ID, SEGUIDOR_ID) VALUES (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

/*Unfollow permite que um usuário deixe de seguir outro*/
func (repositorio usuarios) Unfollow(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM SEGUIDORES WHERE USUARIO_ID = ? AND SEGUIDOR_ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

/*SearchFollowers tem como objetivo buscar os seguidores*/
func (repositorio usuarios) SearchFollowers(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT USUARIOS.ID, USUARIOS.NOME, USUARIOS.NICK, USUARIOS.EMAIL, USUARIOS.CRIADOEM
	  FROM USUARIOS 
	 INNER JOIN SEGUIDORES 
	    ON USUARIOS.ID = SEGUIDORES.SEGUIDOR_ID
	 WHERE SEGUIDORES.USUARIO_ID = ?	
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

/*SearchFollowing traz todos os usuários que um determinado usuário está seguindo*/
func (repositorio usuarios) SearchFollowing(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT USUARIOS.ID, USUARIOS.NOME, USUARIOS.NICK, USUARIOS.EMAIL, USUARIOS.CRIADOEM
	  FROM USUARIOS 
	 INNER JOIN SEGUIDORES 
	    ON USUARIOS.ID = SEGUIDORES.USUARIO_ID
	 WHERE SEGUIDORES.SEGUIDOR_ID = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

/*BuscarSenha traz a senha de um usuario por ID*/
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT SENHA FROM USUARIOS WHERE ID = ?", usuarioID)
	/*Verificar se existe algum erro*/
	if erro != nil {
		/*Se existir retorna um usuário vazio, e o erro*/
		return "", erro
	}
	/*fechar a linha*/
	defer linha.Close()

	var usuario modelos.Usuario
	/*iterar a variável linha*/
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil

}

/*UpdatePassword serve para atualização de senha do usuário*/
func (repositorio usuarios) UpdatePassword(usuariodID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE USUARIOS SET SENHA = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuariodID); erro != nil {
		return erro
	}

	return nil
}
