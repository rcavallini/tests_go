package main

import (
	"fmt"
	"time"
)

func LargadaCorrida() {
	sayHello()
	go corrida("Ferrari")
	go corrida("Mercedes")
	go corrida("RBR")
	go corrida("BMW")
	go corrida("Renault")

	//var input string
	//fmt.Scanln(&input)
	time.Sleep(35 * time.Second)
}

func corrida(carro string) {
	/*
		Essa função vai simular o comportamento de uma carro em uma corrida de 30 voltas.
		Vamos usar a função Sleep para gerar um "delay" de 1 segundo entre os prints.
		Esses delays também serão utilizados pelo escalonador da linguagem como
		critério de interrupção para alternancia entre as rotinas.
	*/
	for i := 0; i <= 30; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(fmt.Sprintf("%s : %d voltas", carro, i))
	}
}
