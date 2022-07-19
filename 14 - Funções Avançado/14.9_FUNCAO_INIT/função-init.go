package main

import "fmt"

var n int

func init() {
	fmt.Println("Função Init. Esta função sempre será executada antes da função main")
	n = 10
}

func main() {
	fmt.Println("Função main")
	//Imprimindo variável n setada na função init
	fmt.Println(n)
}
