package main

import (
	"fmt"
	"io"
	"os"
)

//CountDown is a timer "3 2 1 Go!"
func CountDown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	CountDown(os.Stdout)
}
