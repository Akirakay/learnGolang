package di

import (
	"bytes"
	"fmt"
)

//Greet to the given name by using writer
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprint(writer, "Hello, "+name)
}
