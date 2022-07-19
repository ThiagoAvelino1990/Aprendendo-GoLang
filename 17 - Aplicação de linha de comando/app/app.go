//pacote app que contém a aplicação linha de comando
package app

import (
	"fmt"
	"log"
	"net"

	//Pacote que foi baixado do github
	"github.com/urfave/cli"
)

//Função com a primeira letra maíuscula para o pacote main importar essa função. Necessário comentar oque a função faz
//Função Gerar vai retornar a aplicação de linha de comando pronta para ser executada

func Gerar() *cli.App {
	//Variável app do tipo ponteiro de APP
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IP'S e nomes de Servidor na internet"

	//variável flags do tipo slice cli.Flag []cli.Flag
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br", //Valor padrão passado para a flag, caso a mesma não seja específicada
		},
	}

	//slice de comandos. É um slice pois, podemos ter mais de um comando na aplicação
	//abaixo o tipo Command é um struc
	app.Commands = []cli.Command{
		{
			Name:   "ip",                                 //Propriedade Name que seria o nome do comando
			Usage:  "Busca Ip's de endereço na internet", //Propriedade Usage que explica para que serve o comando
			Flags:  flags,                                //Flags são parâmetros que serão passados para o comando funcionar
			Action: buscarIps,                            //Configuração do que o comando irá fazer. Ação que irá utilizar a função buscarIps
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	//retorno da função Gerar
	return app
}

//Função que passa o parâmetro ponteiro cli.Context
func buscarIps(c *cli.Context) {
	//Buscando o valor da Flag "host" e colocar na variável host
	host := c.String("host")

	//Função do pacote net
	//Irá retornar um slice de ip's e um erro
	//Pode ser um slice pois devido ao redirecionamento dos sites ele pode ter mais de um ip público
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	//_ utilização para não mostrar o índice
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // NS - name server
	if erro != nil {
		log.Fatal(erro)
	}

	//iterar por cada um dos servodres
	for _, servidor := range servidores {
		//
		fmt.Println(servidor.Host)
	}
}
