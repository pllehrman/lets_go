// Source: https://quii.gitbook.io/learn-go-with-tests

package main

import (
	"fmt"
)

// Constants can be good in terms of capturing values and aiding performance.
const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// This line could be rewritten as return prefix
	return 
}

func main() {
	fmt.Println(Hello("world", "English"))
}