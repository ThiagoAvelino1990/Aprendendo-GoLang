package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template" /*Pacote para importação de arquivo HTML e deixar o mesmo mais dinâmico*/
)

//Declaração de variável do tipo ponteiro de template
var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {

	u := usuario{"Joao", "joao@hotmail.com"}

	//Execução do template específico que está dentro de uma variável do tipo template
	//1º Parâmetro - ResponseWriter
	//2º Parâmetro - Nome do arquivo e extensão html
	//3º Parâmetro - Dados a serem mandados para o arquivo html
	templates.ExecuteTemplate(w, "home.html", u)
}

//Para cada campo do struct, deve-se começar com a letra maíuscula, senão não irá funcionar
type usuario struct {
	Nome  string
	Email string
}

func main() {

	//Espressão para referenciar o tipo de arquivo, no caso todos arquivos como html
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Rodando...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
