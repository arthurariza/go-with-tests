package main

import "fmt"

const helloEnglishPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloEnglishPrefix + name
}
