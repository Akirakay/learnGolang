package main

import "fmt"

func hello(name string) string {
	return "hello golang -- from " + name
}

func main() {
	fmt.Println(hello("Akira"))
}
