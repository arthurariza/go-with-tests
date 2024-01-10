package main

import "fmt"

const (
	helloEnglishPrefix = "Hello, "
	helloSpanishPrefix = "Hola, "
	helloFrenchPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Chris", "English"))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloEnglishPrefix

	switch language {
	case "Spanish":
		prefix = helloSpanishPrefix
	case "French":
		prefix = helloFrenchPrefix
	}

	return prefix + name
}
