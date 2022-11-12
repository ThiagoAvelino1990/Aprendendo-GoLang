package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

/*Hash recebe uma string e coloca um hash na senha*/
func Hash(senha string) ([]byte, error) {
	/*pacote bcrypt
	utilização da função GenerateFromPassword
	recebe um slice de byte e um tamanho padrão do pacote bcrypt*/
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

/*VerificarSenha compara uma senha e um hash e retorna se elas são iguais*/
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))

}
