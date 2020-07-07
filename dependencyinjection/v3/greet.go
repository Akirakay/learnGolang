package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

//MyGreeterHandler greet from web
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "akira")
}

//Greet to the given name by using writer
func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, "Hello, "+name)
}
