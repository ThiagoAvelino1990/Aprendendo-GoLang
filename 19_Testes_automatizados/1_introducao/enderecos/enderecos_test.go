/*Para realizar testes das funções o arquivo de teste é necessário ser "nome_test.go"
Desta maneira o GO no momento de realizar o processo de compilação do seu programa, irá verificar o arquivo de nome test.go
 e irá utilizar os mesmos para realizar os teste
TESTE DE UNIDADE - é um teste que ira testar a menor unidade do seu código.
Testes são a única exeção aonde se pode ter pacotes diferentes na mesma pasta
*/
//Pode-se mudar o pacote informando _test
package enderecos_test

import (
	//Para a importação do pacote informar o ponto(.) como um alias
	. "introducao-testes/enderecos"
	"testing"
)

//Refaturando o arquivo de teste
//Criando um Struct
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//Sintaxe de uma função de teste.
//Começa obrigatóriamente com a palavra Test e em seguida o nome da função com a letra maísucula
//Esta é uma boa prática, mas é obritório a função começar com a palavra Test
//Função recebe um parâmetro específico que seria um ponteiro do tipo T
func TestTipoDeEndereco(t *testing.T) {
	//Método para rodar a função em paralelo cou outra
	t.Parallel()

	//Criando um slice de struct
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça Floriano Peixoto", "Praça"},
		{"", "Tipo Inválido"},
		{"RUA DO GABIRU", "Rua"},
		{"ESTRADA DO ALVARENGA", "Estrada"},
		{"rOdOviA 5", "Rodovia"},
	}

	//for para iterar todos os dados contidos no struct
	for _, cenario := range cenariosDeTeste {
		//Recebe o retorno da função de acordo com o endereco inserido do struct
		retornoTipoendereco := TipoDeEndereco(cenario.enderecoInserido)
		//Valida se o retorno da função é igual ao campo do struct "retornoEsperado"
		if retornoTipoendereco != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoTipoendereco,
				cenario.retornoEsperado)
		}
	}
	/*********
	//Linhas comentadas, função refaturada
	//Variável string que recebe o valor - Declaração por inferência de tipo
	enderecoParaTeste := "Avenida Paulista"

	//Variável string que recebe o valor - Declaração por inferência de tipo
	tipoDeEnderecoEsperado := "Avenida"

	//Variavel que recebe o retorno da função TipoDeEndereco
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		//Método Error chamado
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e  recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}
	*********/
}

//Função criada para testar o comando go test -v
func TestQualquer(t *testing.T) {
	//Método para rodar a função em paralelo cou outra
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
