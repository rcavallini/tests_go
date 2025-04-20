package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := ""
	prefix = englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello("Renan", "French"))
	fmt.Println(Hello("Renan", "Spanish"))
	fmt.Println(Hello("Renan", ""))
	fmt.Println(Hello("", ""))
}
