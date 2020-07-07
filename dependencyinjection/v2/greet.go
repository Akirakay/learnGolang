package di

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "akira")
}

//Greet to the given name by using writer
func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, "Hello, "+name)
}
