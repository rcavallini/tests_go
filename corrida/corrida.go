package main

import (
	"fmt"
	"strings"
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
		time.Sleep(1 * time.Second)
		// fmt.Println(carro, ":", i, "voltas")
		fmt.Printf("%s : %d voltas\n", carro, i)
	}
}

func main() {
	fmt.Println("Preparados! Foi dada a largada!")
	var voltas int
	fmt.Println("Quantas voltas a corrida terá?")
	fmt.Scanln(&voltas)
	LargadaCorrida(voltas) // Alterar para passar um numero inserido pelo usuario
	fmt.Println("Corrida finalizada!")
	var escuderia string
	fmt.Println("Qual foi a escuderia vencedora?")
	fmt.Scanln(&escuderia)

	classificacao := map[string]int{
		"FERRARI":  10,
		"MERCEDES": 9,
		"RBR":      6,
		"BMW":      4,
		"RENAULT":  4,
	}

	// fmt.Println(vencedores)

	if value, ok := classificacao[strings.ToUpper(escuderia)]; ok {
		fmt.Printf("A escuderia %s venceu com %d voltas\n", strings.ToUpper(escuderia), value)
		fmt.Println("Total de corridas vencidas:", strings.ToUpper(escuderia), value, ok)

	} else {
		fmt.Println("Escuderia não encontrada")
	}

}
