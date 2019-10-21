package main

import (
	"fmt"
	"os"
	"strings"
)

// Hello says hello in language you specified; English, Spanish or French
func Hello(name, language string) string {

	if name == "" {
		return "Hello, World"
	}
	return greetingPrefix(language) + name

}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Exiting..")
		os.Exit(1)
	}
	name, language := os.Args[1], os.Args[2]
	Hello(name, language)
}

func greetingPrefix(language string) (prefix string) {
	const spanishPrefix = "Hola, "
	const englishPrefix = "Hello, "
	const frenchPrefix = "Bonjour, "
	const spanish = "spanish"
	const english = "english"
	const french = "french"

	prefix = englishPrefix

	switch strings.ToLower(language) {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}
	return prefix

}

