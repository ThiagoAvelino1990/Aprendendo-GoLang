package modelos

import (
	"api/src/seguranca"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/pkg/errors"
)

/*Struct representa um usuário utilizando a rede social*/
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` /*omitempty é importante caso é passado algum dado para o json e que não tenha valor não irá enviar este campo*/
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

/**********************
 *Métodos de validação*
 **********************/
/*Preparar, irá receber os métodos para validar e formatar usuários criados*/
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

/*Validar se os campos obrigatórios foram preenchidos*/
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome deve ser Informado!")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick deve ser Informado!")
	}
	if usuario.Email == "" {
		return errors.New("O Email deve ser Informado!")
	}
	/*Pacote valida se o formado de email é válido e não se existe*/
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("Email inserido inválido !")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha deve ser Informada!")
	}

	return nil
}

/*Retirar espaços em branco de Nome, Nick, Email*/
func (usuario *Usuario) formatar(etapa string) error {
	/*TrimSpace -> utilizado para remover os espaços*/
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
