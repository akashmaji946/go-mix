package main

import (
	"os"

	"github.com/akashmaji946/go-mix/repl"
)

func main() {
	// fmt.Println("Hello, go-mix!")

	// This will only work for arithmetic, bitwise, boolean expressions
	// For now, it works only with binary and boolean expressions involving literals
	repler := repl.Repl{}
	repler.Start(os.Stdin, os.Stdout)
}
