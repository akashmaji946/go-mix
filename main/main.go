package main

import (
	"os"

	"github.com/akashmaji946/go-mix/eval"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/repl"
	"github.com/fatih/color"
)

var MODE = "repl"

var VERSION = "v0.1"
var AUTHOR = "akashmaji(@iisc.ac.in)"
var LICENCE = "MIT"
var PROMPT = "GO-MIX >>> "
var BANNER = `                                                        
    ▄▄▄▄                       ▄▄▄  ▄▄▄     ██              
  ██▀▀▀▀█                      ███  ███     ▀▀              
 ██         ▄████▄             ████████   ████     ▀██  ██▀ 
 ██  ▄▄▄▄  ██▀  ▀██   	       ██ ██ ██     ██       ████   
 ██  ▀▀██  ██    ██   █████    ██ ▀▀ ██     ██       ▄██▄   
  ██▄▄▄██  ▀██▄▄██▀            ██    ██  ▄▄▄██▄▄▄   ▄█▀▀█▄  
    ▀▀▀▀     ▀▀▀▀              ▀▀    ▀▀  ▀▀▀▀▀▀▀▀  ▀▀▀  ▀▀▀                                                       
`
var LINE = "----------------------------------------------------------------"

// Color definitions for file execution output
var (
	redColor    = color.New(color.FgRed)
	yellowColor = color.New(color.FgYellow)
	cyanColor   = color.New(color.FgCyan)
)

func main() {
	// Check if a file argument is provided
	if len(os.Args) > 1 {
		// File mode: read and run a file
		runFile(os.Args[1])
	} else {
		// REPL mode
		repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE, PROMPT)
		repler.Start(os.Stdin, os.Stdout)
	}
}

// runFile reads and executes a Go-Mix source file
func runFile(fileName string) {
	// Read the file contents
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		redColor.Fprintf(os.Stderr, "[FILE ERROR] Could not read file '%s': %v\n", fileName, err)
		os.Exit(1)
	}

	// Convert file contents to string
	source := string(fileContent)

	// Display file info
	// cyanColor.Fprintf(os.Stdout, "Running: %s\n", fileName)
	// fmt.Println(LINE)

	// Execute with panic recovery
	executeFileWithRecovery(source)
}

// executeFileWithRecovery handles parsing and evaluation with panic recovery
func executeFileWithRecovery(source string) {
	// Recover from any panics that might occur during parsing or evaluation
	defer func() {
		if recovered := recover(); recovered != nil {
			redColor.Fprintf(os.Stderr, "[RUNTIME ERROR] %v\n", recovered)
			os.Exit(1)
		}
	}()

	// Parse the source code
	par := parser.NewParser(source)
	rootNode := par.Parse()

	// Check for parser errors
	if par.HasErrors() {
		for _, err := range par.GetErrors() {
			redColor.Fprintf(os.Stderr, "[PARSE ERROR] %s\n", err)
		}
		os.Exit(1)
	}

	if rootNode == nil {
		redColor.Fprintf(os.Stderr, "[PARSE ERROR] Invalid syntax or parser error\n")
		os.Exit(1)
	}

	// Create evaluator and execute
	evaluator := eval.NewEvaluator()
	evaluator.SetParser(par)
	result := evaluator.Eval(rootNode)

	// Display result if any (and not nil)
	if result != nil && result.GetType() != "nil" {
		if result.GetType() == "error" {
			redColor.Fprintf(os.Stderr, "[ERROR] %s\n", result.ToString())
			os.Exit(1)
		} else {
			yellowColor.Fprintf(os.Stdout, "%s\n", result.ToString())
		}
	}
}
