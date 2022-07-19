//Concorrencia é um dos pontos principais em GO. É fácil de utilizar na linguagem e pode melhorar a performance do programa
//Concorrência != Paralelismo
//Paralelismo acontece quando se tem duas ou mais tarefas executas ao mesmo tempo. Só é possível com procesador com mais de um núcleo. A tarefas são distribuídas entre os núcleos
//Tarefas que disputam de forma concorrente, não necessariamente são executadas ao mesmo tempo.
//Caso tenha processador com um núcleo só é possível aplicar concorrência nele
//A diferença é que uma não aguardaria a outra acabar para rodar. De certa forma, as tarefas revezariam o tempo
//GOROUTINE são funções ou métodos em que você pode chamar a execução deles e não necessariamente esperar que elas terminnem para continuar com o programa
package main

import (
	"fmt"
	"time"
)

func main() {
	//informando a palavra go antes de uma função para demarcar que a mesma seja uma goroutine
	go escrever("Olá Mundo!") //goroutine
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
