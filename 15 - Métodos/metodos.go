package main

import "fmt"

//Criação do struct
type usuario struct {
	nome  string
	idade uint8
}

//Declarando um método
//A declaração é iniciada com a cláusula func, mas passando alguma amarração, no caso uma struct
//A sintaxe após informar a amação, é parecida com uma função. Aonde passamos um nome para o método, parâmetros se ele tiver, e um retorno também
func (u usuario) salvar() {
	fmt.Println("Dentro do método Salvar")
	//Printf para imprimir string, o comando com o percentual(%) seguido de uma letra, determina o tipo de dado a ser impresso.
	// Além disso, ao final deve-se informar a vírgula e o tipo de dado a ser impresso
	fmt.Printf("Salvando dados do usuário %s no banco de dados\n", u.nome)
}

//Criação de um método com retorno
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Criação de um método que irá alterar os dados
//Métodos que vão alterar os dados, a amarração deve ser passada como ponteiro pois,
// neste caso, não está sendo feito a cópia dos dados, e sim a alteração do mesmo
//o retorno não precisa realizar o processo de desreferenciação
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("----------MÉTODOS----------")

	//Criação da variável usuario1 do tipo struct.
	//Lembrando que ao criar uma variável do tipo struct, na forma implícita, deve-se se passar os valores dos campos que contém na struct
	usuario1 := usuario{"User 1", 20}
	fmt.Println(usuario1)
	//Chamando o método salvar() que está amarrado a struct usuário e que a variável usuário1 é do tipo usuario(struct)
	usuario1.salvar()

	//Criação da variável usuario2 do tipo struct
	usuario2 := usuario{"Davi", 30}
	fmt.Println(usuario2)
	usuario2.salvar()
	//Retorando o valor do método maiorDeIdade
	fmt.Println(usuario2.maiorDeIdade())

	//chamando o método que utiliza ponteiro.
	usuario1.fazerAniversario()
	//Validando a alteração do dado, após execução do método com ponteiro
	fmt.Println(usuario1.idade)

}
