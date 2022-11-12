package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

/*Publicacao representa um repositório de publicações*/
type Publicacoes struct {
	db *sql.DB
}

/*NovoRepositorioDePublicacoes cria um repositorio de publicacoes*/
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}

}

/*Criar insere uma publicacao no banco de dados*/
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO PUBLICACOES (TITULO, CONTEUDO, AUTOR_ID) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	/*Pegar o último ID inserido*/
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

/*BuscarPorID traz uma única publicação do banco de dados*/
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT PUBLICACOES.*, USUARIOS.NICK
		  FROM PUBLICACOES, USUARIOS
		 WHERE PUBLICACOES.AUTOR_ID = USUARIOS.ID
		   AND PUBLICACOES.ID = ?
	`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.Autornick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil

}

/*buscar traz as publicações dos usuáripos seguidos e também do próprio usuário que fez a requisição*/
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT PUBLICACOES.ID, PUBLICACOES.TITULO, PUBLICACOES.CONTEUDO, PUBLICACOES.AUTOR_ID, PUBLICACOES.CURTIDAS, PUBLICACOES.CRIADAEM, USUARIOS.NICK
	  FROM PUBLICACOES, USUARIOS, SEGUIDORES
	 WHERE USUARIOS.ID = PUBLICACOES.AUTOR_ID
	   AND PUBLICACOES.AUTOR_ID = SEGUIDORES.USUARIO_ID
	   AND (USUARIOS.ID = ? OR SEGUIDORES.SEGUIDOR_ID = ?)
	 GROUP BY PUBLICACOES.ID, PUBLICACOES.TITULO, PUBLICACOES.CONTEUDO, PUBLICACOES.AUTOR_ID, PUBLICACOES.CURTIDAS, PUBLICACOES.CRIADAEM, USUARIOS.NICK
	 ORDER BY 1 DESC
`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.Autornick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

/*Atualizar tem como objetivo alterar os dados de uma publicação no banco de dados*/
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE PUBLICACOES SET TITULO = ?, CONTEUDO = ?
    WHERE ID = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(&publicacao.Titulo, &publicacao.Conteudo, &publicacao.ID); erro != nil {
		return erro
	}

	return nil

}

/*Deletar exclui uma publicação no banco de dados*/
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
	DELETE FROM PUBLICACOES WHERE ID = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(&publicacaoID); erro != nil {
		return erro
	}

	return nil
}

/*BuscarPorUsuario busca publicacoes por usuario*/
func (repositorio Publicacoes) BuscarPorUsuario(usuariodID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT PUBLICACOES.ID, PUBLICACOES.TITULO, PUBLICACOES.CONTEUDO, PUBLICACOES.AUTOR_ID, PUBLICACOES.CURTIDAS, PUBLICACOES.CRIADAEM, USUARIOS.NICK
		  FROM PUBLICACOES, USUARIOS
		 WHERE PUBLICACOES.AUTOR_ID = USUARIOS.ID
		   AND USUARIOS.ID = ?`,
		usuariodID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.Autornick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

/*Curtir adiciona uma curtida na publicacao*/
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE PUBLICACOES SET CURTIDAS = CURTIDAS + 1
    	WHERE  ID = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(&publicacaoID); erro != nil {
		return erro
	}

	return nil
}

/*Descurtir adiciona uma curtida na publicacao*/
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE PUBLICACOES SET CURTIDAS = 
			CASE
				WHEN CURTIDAS > 0 THEN CURTIDAS -1
				ELSE 0
			END
		WHERE ID = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(&publicacaoID); erro != nil {
		return erro
	}

	return nil
}
