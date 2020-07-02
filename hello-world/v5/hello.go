package main

import "fmt"

const englishHelloPrefix = "hello golang -- from "

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed.
func Hello(name string) string {
	if name == "" {
		name = "Akira"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Akira"))
}
