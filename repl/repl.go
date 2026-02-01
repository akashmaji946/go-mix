package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/akashmaji946/go-mix/eval"
	"github.com/akashmaji946/go-mix/parser"
)

type Repl struct {
}

func NewRepl() *Repl {
	return &Repl{}
}

func (r *Repl) Start(reader io.Reader, writer io.Writer) {
	fmt.Println("Welcome to go-mix!")
	fmt.Println("Type your code and press enter")
	fmt.Println("Type '.exit' to quit")
	fmt.Println("")

	scanner := bufio.NewScanner(reader)
	evaluator := eval.NewEvaluator()
	for {
		fmt.Print(">>> ")
		scanned := scanner.Scan()
		if !scanned {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		line := scanner.Text()
		if line == ".exit" {
			writer.Write([]byte("Good Bye!\n"))
			break
		}
		parser := parser.NewParser(line)
		rootNode := parser.Parse()
		if rootNode == nil {
			writer.Write([]byte("Invalid syntax or parser error\n"))
			continue
		}

		result := evaluator.Eval(rootNode)

		if result != nil {
			writer.Write([]byte(result.ToString() + "\n"))
		}

	}

}
