/*
File    : go-mix/main/main.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)

Package main is the entry point for the Go-Mix interpreter.
It provides two modes of operation:
1. REPL Mode (default): Interactive Read-Eval-Print Loop for live coding
2. File Mode: Execute Go-Mix source files from the command line

The interpreter uses a lexer-parser-evaluator pipeline to process Go-Mix code.
*/
package main

import (
	"os"

	"github.com/akashmaji946/go-mix/eval"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/repl"
	"github.com/fatih/color"
)

// MODE defines the default operating mode of the interpreter
// Currently set to "repl" for interactive mode
var MODE = "repl"

// VERSION represents the current version of the Go-Mix interpreter
var VERSION = "v0.1"

// AUTHOR contains the contact information of the interpreter's author
var AUTHOR = "akashmaji(@iisc.ac.in)"

// LICENCE specifies the software license (MIT License)
var LICENCE = "MIT"

// PROMPT is the command prompt displayed in REPL mode
var PROMPT = "Go-Mix >>> "

// BANNER is the ASCII art logo displayed when starting the REPL
// It shows "Go-Mix" in stylized ASCII characters
var BANNER = `                                                        
    ▄▄▄▄                       ▄▄▄  ▄▄▄     ██              
  ██▀▀▀▀█                      ███  ███     ▀▀              
 ██         ▄████▄             ████████   ████     ▀██  ██▀ 
 ██  ▄▄▄▄  ██▀  ▀██   	       ██ ██ ██     ██       ████   
 ██  ▀▀██  ██    ██   █████    ██ ▀▀ ██     ██       ▄██▄   
  ██▄▄▄██  ▀██▄▄██▀            ██    ██  ▄▄▄██▄▄▄   ▄█▀▀█▄  
    ▀▀▀▀     ▀▀▀▀              ▀▀    ▀▀  ▀▀▀▀▀▀▀▀  ▀▀▀  ▀▀▀                                                       
`

// LINE is a separator line used for visual formatting in the REPL
var LINE = "----------------------------------------------------------------"

// Color definitions for file execution output
// These colors are used to provide visual feedback during file execution:
// - redColor: Error messages and critical failures
// - yellowColor: Normal output and results
// - cyanColor: Informational messages
var (
	redColor    = color.New(color.FgRed)
	yellowColor = color.New(color.FgYellow)
	cyanColor   = color.New(color.FgCyan)
)

// main is the entry point of the Go-Mix interpreter.
// It determines the operating mode based on command-line arguments:
//
// Usage:
//
//	go-mix              - Start in REPL (interactive) mode
//	go-mix <filename>   - Execute the specified Go-Mix source file
//
// The function delegates to either runFile() for file execution
// or starts the REPL for interactive programming.
func main() {
	// Check if a file argument is provided
	// os.Args[0] is the program name, os.Args[1] would be the first argument
	if len(os.Args) > 1 {
		// File mode: read and run a file
		fileName := os.Args[1]
		runFile(fileName)
	} else {
		// REPL mode: Start interactive interpreter
		// Create a new REPL instance with banner, version info, and prompt
		repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE, PROMPT)
		// Start the REPL loop, reading from stdin and writing to stdout
		repler.Start(os.Stdin, os.Stdout)
	}
}

// runFile reads and executes a Go-Mix source file.
// It handles the complete file execution pipeline:
// 1. Read the file from disk
// 2. Convert contents to string
// 3. Execute the code with error recovery
//
// Parameters:
//
//	fileName - Path to the Go-Mix source file to execute
//
// Error Handling:
//   - File read errors: Displays error message and exits with code 1
//   - Parse/runtime errors: Handled by executeFileWithRecovery()
func runFile(fileName string) {
	// Read the file contents using os.ReadFile (replaces deprecated ioutil.ReadFile)
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		// Display file read error in red and exit
		redColor.Fprintf(os.Stderr, "[FILE ERROR] Could not read file '%s': %v\n", fileName, err)
		os.Exit(1)
	}

	// Convert file contents from []byte to string for parsing
	source := string(fileContent)

	// Display file info (currently commented out for cleaner output)
	// cyanColor.Fprintf(os.Stdout, "Running: %s\n", fileName)
	// fmt.Println(LINE)

	// Execute the source code with panic recovery to handle runtime errors gracefully
	executeFileWithRecovery(source)
}

// executeFileWithRecovery handles parsing and evaluation with panic recovery.
// This function implements a robust error handling strategy:
// 1. Sets up panic recovery to catch runtime errors
// 2. Parses the source code into an AST
// 3. Checks for parsing errors
// 4. Evaluates the AST
// 5. Displays results or errors
//
// Parameters:
//
//	source - The Go-Mix source code as a string
//
// Error Handling:
//   - Panics: Caught by defer/recover, displayed as runtime errors
//   - Parse errors: Collected and displayed, then exit
//   - Evaluation errors: Displayed in red, then exit
//   - Success: Result displayed in yellow (if not nil)
func executeFileWithRecovery(source string) {
	// Recover from any panics that might occur during parsing or evaluation
	// This prevents the interpreter from crashing and provides user-friendly error messages
	defer func() {
		if recovered := recover(); recovered != nil {
			redColor.Fprintf(os.Stderr, "[RUNTIME ERROR] %v\n", recovered)
			os.Exit(1)
		}
	}()

	// Parse the source code into an Abstract Syntax Tree (AST)
	// The parser performs lexical analysis and syntactic analysis
	par := parser.NewParser(source)
	rootNode := par.Parse()

	// Check for parser errors
	// The parser collects errors instead of panicking, allowing multiple errors to be reported
	if par.HasErrors() {
		for _, err := range par.GetErrors() {
			redColor.Fprintf(os.Stderr, "[PARSE ERROR] %s\n", err)
		}
		os.Exit(1)
	}

	// Verify that parsing produced a valid AST root node
	if rootNode == nil {
		redColor.Fprintf(os.Stderr, "[PARSE ERROR] Invalid syntax or parser error\n")
		os.Exit(1)
	}

	// Create evaluator and execute the AST
	// The evaluator walks the AST and executes the program
	evaluator := eval.NewEvaluator()
	evaluator.SetParser(par) // Link parser for access to environment and error handling
	result := evaluator.Eval(rootNode)

	// Display result if any (and not nil)
	// Only display non-nil results to avoid cluttering output
	if result != nil && result.GetType() != "nil" {
		if result.GetType() == "error" {
			// Evaluation produced an error object - display and exit
			redColor.Fprintf(os.Stderr, "[ERROR] %s\n", result.ToString())
			os.Exit(1)
		} else {
			// Successful evaluation - display result in yellow
			yellowColor.Fprintf(os.Stdout, "%s\n", result.ToString())
		}
	}
}
