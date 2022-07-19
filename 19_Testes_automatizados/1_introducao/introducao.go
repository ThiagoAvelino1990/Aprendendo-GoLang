package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	fmt.Println("Testes Automarizados")
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia do Imigrantes")
	fmt.Println(tipoEndereco)

}
