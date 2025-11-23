package main

import (
	"os"

	"github.com/danilobml/monkey-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
