package enderecos

import "strings"

//TipoDeEndereço verifica o endereço digitado e retornar qual o tipo de endereço
func TipoDeEndereco(end string) string {
	//Criando um slice de string com os tipos de endereços
	tiposValidos := []string{"rua", "avenida", "praça", "estrada", "rodovia"}

	//Passando o valor digitado na função para letras minusculas
	enderecoEmLetraMinuscula := strings.ToLower(end)

	//Função Split dentro do pacote strings que basicamente divide uma string em um slice de acordo com o separador passado
	//No caso abaixo, o separador é o espaço em branco " ".
	//O valor[0] significa que eu quero apenas a primeira posição deste slice
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	//Variável boolean
	enderecoTemUmTipoValido := false

	//For sem contador
	//a função range na declaração da variável faz com que a varíavel valide todos os tipos descritos no slice
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	//Retornando e transformando o retorno da função na primeira letra maíuscula se enderecoTemUmTipoValido for true
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}
