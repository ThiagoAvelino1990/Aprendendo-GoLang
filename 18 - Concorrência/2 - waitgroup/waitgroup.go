//waitGroup é uma das formas existentes para sincronizar a execução das goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Variável do tipo sync.Waitgroup
	var waitGroup sync.WaitGroup

	//Utilizando a propriedade de quantas goroutines existem, que irão rodar ao mesmo tempo. Deve-se passar a quantidade
	waitGroup.Add(4)

	//Função anônima com gouroutine
	go func() {
		escrever("Olá Mundo!")
		//método basicamente faz -1 no contador acima. waitGroup.add()
		waitGroup.Done()

	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Gouroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Gouroutine 4!")
		waitGroup.Done()
	}()

	//Aguarda a finalização da execução das funções goroutine
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
