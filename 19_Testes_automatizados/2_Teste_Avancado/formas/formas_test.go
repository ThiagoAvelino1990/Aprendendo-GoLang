//Nomenclatura indicando que se faz referencia ao um arquivo test
package formas_test

import (
	"math"
	//Para a importação do pacote informar o ponto(.) como um alias
	. "teste-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	//TDD - Test Driven Development(Desenvolvimento guiado a testes)
	//
	//t.Run é usado quando se quer separar as coisas.
	t.Run("Retângulo", func(t *testing.T) {
		//variável que recebe os valores do struc Retangulo/formas.go
		ret := Retangulo{10, 12}

		areaEsperada := float64(120)
		//Chamando o método Area() referenciado ao struct Retangulo/formas.go
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			//t.Fatalf faz o mesmo que o t.Errorf, porém ele irá parar a execução
			//o Errorf vai dar erro mas vai continuar a execução
			t.Fatalf("A área recebida do Retângulo %f é difernete da esperada %f",
				areaRecebida, areaEsperada)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida do Circulo %f é difernete da esperada %f",
				areaRecebida, areaEsperada)
		}

	})

	t.Run("Triângulo", func(t *testing.T) {
		tri := Triangulo{5, 10}

		areaEsperada := float64(25)
		areaRecebida := tri.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida do Triângulo %f é difernete da esperada %f",
				areaRecebida, areaEsperada)
		}

	})
}
