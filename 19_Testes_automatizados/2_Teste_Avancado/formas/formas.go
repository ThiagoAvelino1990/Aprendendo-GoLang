//Os métodos, interfaces, métodos devem ter comentários e com a letra maíscula para a importação
package formas

import (
	"math"
)

//Criando struct retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

//Criando struct circulo
type Circulo struct {
	Raio float64
}

//Criando struct triângulo
type Triangulo struct {
	Base   float64
	Altura float64
}

//Criação de uma interface(Deve conter ao menos um método)
type Forma interface {
	//Método chamado area que não recebe parâmetros e retornar um float64
	Area() float64
}

/***********MÉTODOS**************/
//Criação do método area para a forma retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

//Criação do método area para a forma circulo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

//Criação do método area para a forma triangulo
func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}
