package main

import "fmt"

const englishHelloPrefix = "hello golang -- from "

func hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Akira"))
}
