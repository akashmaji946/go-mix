package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/akashmaji946/go-mix/eval"
	"github.com/akashmaji946/go-mix/parser"
)

type Repl struct {
	Banner  string
	Version string
	Author  string
	Line    string
	License string
}

func NewRepl(banner string, version string, author string, line string, license string) *Repl {
	return &Repl{Banner: banner, Version: version, Author: author, Line: line, License: license}
}

func (r *Repl) Start(reader io.Reader, writer io.Writer) {
	fmt.Println(r.Line)
	fmt.Println(r.Banner)
	fmt.Println(r.Line)
	fmt.Println("Version: " + r.Version + " | Author: " + r.Author + " | Lincense: " + r.License)
	fmt.Println(r.Line)
	fmt.Println("Welcome to Go-Mix!")
	fmt.Println("Type your code and press enter")
	fmt.Println("Type '.exit' to quit")
	fmt.Println(r.Line)

	scanner := bufio.NewScanner(reader)
	evaluator := eval.NewEvaluator()

	for {
		fmt.Print("Go-Mix >>> ")
		scanned := scanner.Scan()
		if !scanned {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		line := scanner.Text()
		if strings.Trim(line, " \n\t\r") == "" {
			continue
		}
		if line == ".exit" {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		// Execute parsing and evaluation with panic recovery
		r.executeWithRecovery(writer, line, evaluator)
	}
}

// executeWithRecovery handles parsing and evaluation with panic recovery
func (r *Repl) executeWithRecovery(writer io.Writer, line string, evaluator *eval.Evaluator) {
	// Recover from any panics that might occur during parsing or evaluation
	defer func() {
		if recovered := recover(); recovered != nil {
			writer.Write([]byte(fmt.Sprintf("[RUNTIME ERROR] %v\n", recovered)))
		}
	}()

	par := parser.NewParser(line)
	rootNode := par.Parse()

	// Check for parser errors
	if par.HasErrors() {
		for _, err := range par.GetErrors() {
			writer.Write([]byte(err + "\n"))
		}
		return
	}

	if rootNode == nil {
		writer.Write([]byte("[LEXER ERROR] Invalid syntax or parser error\n"))
		return
	}

	evaluator.SetParser(par)
	result := evaluator.Eval(rootNode)

	if result != nil {
		writer.Write([]byte(result.ToString() + "\n"))
	}
}
