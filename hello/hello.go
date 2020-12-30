package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Holla "
const frenchHelloPrefix = "Bonjour "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	prefix := englishHelloPrefix

	if name == "" {
		name = "World"
	}

	if language == spanish {
		prefix = spanishHelloPrefix

	}

	if language == french {
		prefix = frenchHelloPrefix

	}

	return prefix + name
}

func main() {

	fmt.Println(Hello("World", "english"))

}
