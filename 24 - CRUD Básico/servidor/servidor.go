package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//Utilização do pacote ioutil - utilizado para entrada e saida de dados
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		//Este return vazio é utilizado apenas para que o programa parece a execução
		return
	}

	/*struct de usuário*/
	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao convereter o usuario para struct!"))
		return
	}

	fmt.Println(usuario)

	/*Abrindo conexão com o banco*/
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close() /*Fechando o banco*/

	/*PREPARE STATEMENT - CRIA COMANDO DE INSERÇÃO UTILIZADO PARA EVITAR O ATAQUE CHAMANDO DE SQLINJECTION*/
	/*criando o comando para inserção de dados*/
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
	}
	defer statement.Close() /*Fechando o statement*/

	/*Realizando a inserção dos dados no banco*/
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	/*Devolvendo o ID do usuário inserido*/
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	/*STATUS CODES - Em resquiçoes HTTP indicam se a requisição deu um problema ou foi um sucesso*/
	/*Neste caso, utilizaremos o created pois, estamos inserindo dados*/
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

//Buscar usuários cadastrados no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	/*Abrindo conexão com o banco de dados*/
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	//SELECT * FROM USUARIOS
	linhas, erro := db.Query("SELECT * FROM USUARIOS")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários no banco!"))
		return
	}
	defer linhas.Close()

	/*Criando um slice de usuario*/
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao convereter os usuários para JSON"))
		return
	}
}

//Buscar usuário específico cadastrado no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("SELECT * FROM USUARIOS WHERE ID = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	/*Conversão para JSON*/
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}

}

/*Atualizar usuário*/
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	/*Lendro parâmetro request*/
	parametros := mux.Vars(r)

	/*Convertendo o parâmetro para inteiro*/
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro !"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição !"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro para converter usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados !"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE USUARIOS SET NOME = ?, EMAIL = ? WHERE ID = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário !"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados !"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM USUARIOS WHERE ID = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement !"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deleter o usuário !"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
