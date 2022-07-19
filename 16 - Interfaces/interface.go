package main

import (
	"fmt"
	"math"
)

//Criando struct retangulo
type retangulo struct {
	altura  float64
	largura float64
}

//Criando struct circulo
type circulo struct {
	raio float64
}

//Criando struct triângulo
type triangulo struct {
	base   float64
	altura float64
}

//Criação de uma interface(Deve conter ao menos um método)
type forma interface {
	//Método chamado area que não recebe parâmetros e retornar um float64
	area() float64
}

//Função para calcular a area
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

/***********MÉTODOS**************/
//Criação do método area para a forma retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

//Criação do método area para a forma circulo
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//Criação do método area para a forma triangulo
func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func main() {
	fmt.Println("--------------INTERFACE--------------")

	r := retangulo{10, 5}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	t := triangulo{5, 10}
	escreverArea(t)

}
