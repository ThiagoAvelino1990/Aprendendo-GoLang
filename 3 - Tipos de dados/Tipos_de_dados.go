package main

import (
	"errors"
	"fmt"
)

func main() {
	//Numeros Inteiros(int) podem ser positivos e negavitos

	//Declaração explícita int8 -> inteiro 8bits
	var inteiro8 int8 = -123
	fmt.Println("Inteiro 8 bits ", inteiro8)

	//Declaração explícita int16 -> inteiro 16 bits
	var inteiro16 int16 = -12345
	fmt.Println("Inteiro 16 bits ", inteiro16)

	//Declaração explícita int32 -> inteiro 32 bits
	var inteiro32 int32 = -1234567890
	fmt.Println("Inteiro 32 bits ", inteiro32)

	//Declaração explícita int64 -> inteiro 64 bits
	var inteiro64 int64 = -1234567890098765432
	fmt.Println("Inteiro 64 bits ", inteiro64)

	//Declaração implícita int -> De acordo com a arquitetura do computador(32bits ou 64bits)
	inteiropc := -1234567890098765432
	fmt.Println("INT - Declaração implícita(Arquitetura computador) ", inteiropc)

	//Numeros Inteiros(uint)
	//unsygned -> Não simbolizados

	//Declaração explícita uint8 -> inteiro 8bits
	var uinteiro8 uint8 = 123
	fmt.Println("Inteiro(unsygned) 8 bits ", uinteiro8)

	//Declaração explícita uint16 -> inteiro 16bits
	var uinteiro16 uint16 = 12345
	fmt.Println("Inteiro(unsygned) 16 bits ", uinteiro16)

	//Declaração explícita uint32-> inteiro 32bits
	var uinteiro32 uint32 = 1234567890
	fmt.Println("Inteiro(unsygned) 32 bits ", uinteiro32)

	//Declaração explícita uint64 -> inteiro 64bits
	var uinteiro64 int64 = 1234567890098765432
	fmt.Println("Inteiro(unsygned) 64 bits ", uinteiro64)

	//Declaração implícita uint -> De acordo com a arquitetura do computador(32bits ou 64bits)
	uinteiropc := 1234567890098765432
	fmt.Println("UINT - Declaração implícita(Arquitetura computador) ", uinteiropc)

	//Alias -> Apelido
	//rune - 32 bits(INT32)
	var alias1 rune = 1234567890
	fmt.Println("Alias(Apelido) RUNE para inteiro 32bits ", alias1)

	//byte - 8bits(UINT8)
	var alias2 byte = 123
	fmt.Println("Alias(Apelido) BYTE para inteiro 8bits ", alias2)

	//Números Reais(Valor deve-se ser com ponto(.) e não com vírgula(,))
	//Declaração explícita float32 -> 32 bits
	var flt32 float32 = 12345.12
	fmt.Println("Declaração explícita float32 ", flt32)

	//Declaração explícita float64 -> 64 bits
	var flt64 float64 = -12345.12
	fmt.Println("Declaração explícita float64 ", flt64)

	//Declaração implícita float -> De acordo com a arquitetura do computador(32bits ou 64bits)
	flt := 1234567890098765432.12
	fmt.Println("Declaração implícita float -> De acordo com a arquitetura do computador(32bits ou 64bits) ", flt)

	//Cadeias de Caracteres
	var str string = "String Explícito"
	fmt.Println(str)

	str2 := "String Implícito"
	fmt.Println(str2)

	var str3 string
	fmt.Println("Declaração com a cláusula VAR e o tipo de dado sem valor: ", str3)

	//CHAR DECLARAÇÃO IMPLÍCITA APENAS
	//Declaração CHAR traz o valor da tabela ASCI
	char := 'B'
	fmt.Println("Declaração CHAR traz o valor da tabela ASCI. char := 'B' ", char)

	//Boolean
	//Declarando tipo bool explícito
	var boolean bool = true
	fmt.Println(boolean)

	var boolean2 bool
	fmt.Println("O Valor 'ZERO' de boolean é: ", boolean2)

	//Declarando o tipo bool Implícito
	boolean3 := true
	fmt.Println(boolean3)

	//ERROR -> Deve-se utilizar o pacote errors
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	//Valor zero para o tipo error
	var erroi error
	fmt.Println("Valor zero para o tipo error ", erroi)

}
