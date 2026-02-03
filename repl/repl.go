package repl

import (
	"io"
	"strings"

	"github.com/akashmaji946/go-mix/eval"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

// Color definitions for REPL output
var (
	blueColor   = color.New(color.FgBlue)
	yellowColor = color.New(color.FgYellow)
	redColor    = color.New(color.FgRed)
	greenColor  = color.New(color.FgGreen)
	cyanColor   = color.New(color.FgCyan)
)

type Repl struct {
	Banner  string
	Version string
	Author  string
	Line    string
	License string
	Prompt  string
}

func NewRepl(banner string, version string, author string, line string, license string, prompt string) *Repl {
	return &Repl{Banner: banner, Version: version, Author: author, Line: line, License: license, Prompt: prompt}
}

func (r *Repl) PrintBannerInfo(writer io.Writer) {

	blueColor.Fprintf(writer, "%s\n", r.Line)
	greenColor.Fprintf(writer, "%s\n", r.Banner)
	blueColor.Fprintf(writer, "%s\n", r.Line)
	yellowColor.Fprintln(writer, "Version: "+r.Version+" | Author: "+r.Author+" | Lincense: "+r.License)
	blueColor.Fprintf(writer, "%s\n", r.Line)

	cyanColor.Fprintf(writer, "%s\n", "Welcome to Go-Mix!")
	cyanColor.Fprintf(writer, "%s\n", "Type your code and press enter")
	cyanColor.Fprintf(writer, "%s\n", "Type '.exit' to quit")
	cyanColor.Fprintf(writer, "%s\n", "Use up/down arrows to navigate command history")
	blueColor.Fprintf(writer, "%s\n", r.Line)
}

func (r *Repl) Start(reader io.Reader, writer io.Writer) {

	// print the banner
	r.PrintBannerInfo(writer)

	// create a new readline instance
	rl, err := readline.New(r.Prompt)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	evaluator := eval.NewEvaluator()

	for {
		line, err := rl.Readline()
		if err != nil {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		line = strings.Trim(line, " \n\t\r")
		if line == "" {
			continue
		}
		if line == ".exit" {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		rl.SaveHistory(line)

		// echo input in blue
		// blueColor.Fprintf(writer, ">> %s\n", line)

		// execute parsing and evaluation with panic recovery
		r.executeWithRecovery(writer, line, evaluator)
	}
}

// executeWithRecovery handles parsing and evaluation with panic recovery
func (r *Repl) executeWithRecovery(writer io.Writer, line string, evaluator *eval.Evaluator) {
	// Recover from any panics that might occur during parsing or evaluation
	defer func() {
		if recovered := recover(); recovered != nil {
			redColor.Fprintf(writer, "[RUNTIME ERROR] %v\n", recovered)
		}
	}()

	par := parser.NewParser(line)
	rootNode := par.Parse()

	// Check for parser errors
	if par.HasErrors() {
		for _, err := range par.GetErrors() {
			redColor.Fprintf(writer, "%s\n", err)
		}
		return
	}

	if rootNode == nil {
		redColor.Fprintf(writer, "[LEXER ERROR] Invalid syntax or parser error\n")
		return
	}

	evaluator.SetParser(par)
	result := evaluator.Eval(rootNode)

	if result != nil {
		if result.GetType() == "error" {
			redColor.Fprintf(writer, "%s\n", result.ToString())
		} else {
			yellowColor.Fprintf(writer, "%s\n", result.ToString())
		}
	}
}
