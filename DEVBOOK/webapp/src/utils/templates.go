package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

/*CarregarTemplates insere os arquivos html na variávle templates*/
func CarregarTemplates() {
	/*sintaxe para bsucar todos os arquivos con extensão html dentro da pasta views*/
	templates = template.Must(template.ParseGlob("views/*.html"))
}

/*ExecutarTemplate renderiza uma página HTML na tela*/
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, nil)
}
