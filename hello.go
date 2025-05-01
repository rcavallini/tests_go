package main

import (
	"errors"
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) (string, error) {
	if name == "" {
		name = "World"
	}

	prefix, err := greetingPrefix(language)
	if err != nil {
		return "", err
	}

	return prefix + name, nil
}

func greetingPrefix(language string) (string, error) {
	switch language {
	case spanish:
		return spanishHelloPrefix, nil
	case french:
		return frenchHelloPrefix, nil
	case english:
		return englishHelloPrefix, nil
	default:
		return "", errors.New("unsupported language: " + language)
	}
}

func sayHello() {
	if msg, err := Hello("", "Spanish"); err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Erro:", err)
	}

	if msg, err := Hello("Renan", "German"); err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Erro:", err)
	}
}
