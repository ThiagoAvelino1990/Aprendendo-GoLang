package modelos

/*Senha representa o forma de requisição de alteração de senha*/
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
