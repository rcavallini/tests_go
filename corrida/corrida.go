package corrida

import (
	"fmt"
	"time"
)

func LargadaCorrida(x int) {
	go corrida("Ferrari", x)
	go corrida("Mercedes", x)
	go corrida("RBR", x)
	go corrida("BMW", x)
	go corrida("Renault", x)

	//var input string
	//fmt.Scanln(&input)
	time.Sleep(time.Duration(x+2) * time.Second)
}

func corrida(carro string, x int) {
	/*
		Essa função vai simular o comportamento de uma carro em uma corrida de 30 voltas.
		Vamos usar a função Sleep para gerar um "delay" de 1 segundo entre os prints.
		Esses delays também serão utilizados pelo escalonador da linguagem como
		critério de interrupção para alternancia entre as rotinas.
	*/
	for i := 0; i <= x; i++ {
		time.Sleep(1000 * time.Millisecond)
		// fmt.Println(carro, ":", i, "voltas")
		fmt.Printf("%s : %d voltas\n", carro, i)
	}
}
