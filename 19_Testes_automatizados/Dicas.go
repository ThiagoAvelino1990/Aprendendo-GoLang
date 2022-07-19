package main

import "fmt"

func main() {
	fmt.Println("45. Dicas do comando go test e relatórios de cobertura")
	fmt.Println("Quando se roda o coamdno 'go test' por padrão ele irá rodar os arquivos que terminam com _test.go.")
	fmt.Println("O comando 'go test' irá executrar todos os testes do pacote ao qual você está.")
	fmt.Println("O comando 'go test ./...' irá fazer com que o GO entre em todos os pacotes do projeto e execute os testes de todos. Caso apareça a descrição (cashed) em um dos arquivos, significa que o GO identificou que não houve alteração de um determinado pacote e isso fará com que ele não execute novamente.")
	fmt.Println("O comando 'go test - v' seria o modo verboso, é algo que irá trazer uma melhor descrição do test realizado")
	fmt.Println("Para rodar as funções em paralelo é necessário informar o comando dentro da função 't.Parallel()'. Este comando deve ser informado na primeira linha da função")
	fmt.Println("Comando para identificar a porcentagem das linhas da função, ou seja, se todas as linhas estão sendo acobertadas pelo test'go test --cover'")
	fmt.Println("Ferramenta para descobrir oque está coberto ou não na função test 'go test --coverprofile nome_do_arquivo.txt")
	fmt.Println("Comando para descobrir quais funções estão sendo testadas e qual delas não foi acobertada em 100%'go tool cover --func=nome_do_arquivo.txt")
	fmt.Println("Comando para mostrar arquivo HTML com todas as linhas que não estão cobertas das funções test 'go tool cover --html=nome_do_arquivo.txt")

}
