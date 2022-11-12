package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go" /*jwt é um alias para o pacote, visto que o mesmo contém um traço no nome final dele*/
)

/*CriarToken retorna um token assinado com as permissões do usuário*/
func CriarToken(usuarioID uint64) (string, error) {
	/*MapClaims são as permissoes que irão ter dentro do token*/
	permissoes := jwt.MapClaims{}
	/*Populando o MAP*/
	permissoes["autorizhed"] = true                          /*autorização para funções da API*/
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() /*"exp" - é uma chave padrão para expirar o tempo do token. Unix() -> refere-se a quantidade de milisegundos*/
	permissoes["usuarioId"] = usuarioID
	/*Gerando uma chave secret
	jwt.SigningMethodHS256 -> Método de assinatura do Token*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) /*Secret KEY*/
}

/*ValidarToken verifica se o token passado na requisição é válido*/
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New(strings.ToUpper("token inválido"))
}

/*extrairToken Esta função tem como objetivo quebrar o token pois,
um token "válido" deve conter duas posições, uma vez que a primeira palavrar
de um token começa por Bearer(portador do token)*/
func extrairToken(r *http.Request) string {
	/*Pegar o token que vem do request*/
	token := r.Header.Get("Authorization")
	/*Valida se o token possui duas strings*/
	if len(strings.Split(token, " ")) == 2 {
		/*retorna a segunda string*/
		return strings.Split(token, " ")[1]
	}

	return ""
}

/*retornarChaveDeVerificacao*/
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf(strings.ToUpper("método de assiantura inesperado %v"), token.Header["alg"])
	}

	return config.SecretKey, nil
}

/*ExtrairUsuarioID retorna o usuarioId que está salvo no token*/
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		/*
			Permissoes["usuarioId"] é do tipo interface e quando passar o valor
			por padrão ficará salvo como float.
			A função ParseUint recebe como primeiro parâmetro um tipo string.
			Deve-se converter o float para string, para o ParseUint transformar a string e um UINT*/
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New(strings.ToUpper("token inválido "))
}
