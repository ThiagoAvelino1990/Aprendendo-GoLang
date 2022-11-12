package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

/*Configurar utiliza as variáveis de ambiente para a criaçaõ do SecureCookie*/
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

/*salvar registra as informações de autenticação*/
func Salvar(w http.ResponseWriter, ID, token string) error {
	/*Criação dos dados pelo MAP*/
	dados := map[string]string{
		"id":     ID,
		"tokren": token,
	}
	/*Codificar os dados*/
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil

}

/*Ler retorna os valores armazenados no cookie*/
func Ler(r *http.Request) (map[string]string, error) {
	/*Realizando a leitura do Cookie decodificado*/
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	/*Alocando um map vazio e jogando na variável valores*/
	valores := make(map[string]string)

	/*Decodificando o Cookie e jogando os dados dentro da variável valores*/
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}
