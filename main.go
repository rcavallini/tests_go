package main

import (
	"fmt"
	"strings"

	"github.com/rcavallini/tests_go/corrida"
)

func main() {
	fmt.Println("Preparados! Foi dada a largada!")
	var voltas int
	fmt.Println("Quantas voltas a corrida terá?")
	fmt.Scanln(&voltas)
	corrida.LargadaCorrida(voltas) // Alterar para passar um numero inserido pelo usuario
	fmt.Println("Corrida finalizada!")
	var escuderia string
	fmt.Println("Qual foi a escuderia vencedora?")
	fmt.Scanln(&escuderia)
	fmt.Println((strings.ToUpper(escuderia)))

}
