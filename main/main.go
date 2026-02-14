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
	"fmt"
	"net"
	"os"

	"github.com/akashmaji946/go-mix/eval"
	_ "github.com/akashmaji946/go-mix/file"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/repl"
	"github.com/fatih/color"
)

// MODE defines the default operating mode of the interpreter
// Currently set to "repl" for interactive mode
var MODE = "repl"

// VERSION represents the current version of the Go-Mix interpreter
var VERSION = "v1.0.0"

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
//	go-mix --help       - Display help information
//	go-mix --version    - Display version information
//
// The function delegates to either runFile() for file execution
// or starts the REPL for interactive programming.
func main() {
	// Check if a flag argument is provided
	if len(os.Args) > 1 {
		arg := os.Args[1]

		// Handle --help flag
		if arg == "--help" || arg == "-h" {
			showHelp()
			os.Exit(0)
		}

		// Handle --version flag
		if arg == "--version" || arg == "-v" {
			showVersion()
			os.Exit(0)
		}

		// Server mode: Start a REPL server
		if arg == "server" {
			if len(os.Args) < 3 {
				redColor.Fprintf(os.Stderr, "[USAGE ERROR] Missing port for server mode. Usage: go-mix server <port>\n")
				os.Exit(1)
			}
			port := os.Args[2]
			startServer(port)
			return // Exit after starting the server
		}
		// File mode: read and run a file
		fileName := arg
		runFile(fileName)
	} else {
		// REPL mode: Start interactive interpreter
		// Create a new REPL instance with banner, version info, and prompt
		repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE, PROMPT)
		// Start the REPL loop, reading from stdin and writing to stdout
		repler.Start(os.Stdin, os.Stdout)
	}
}

// showHelp displays the help information for the Go-Mix interpreter
func showHelp() {
	cyanColor.Println("Go-Mix - An Interpreted Programming Language")
	cyanColor.Println("")
	cyanColor.Println("USAGE:")
	yellowColor.Println("  go-mix                    Start interactive REPL mode")
	yellowColor.Println("  go-mix <path-to-file>     Execute a Go-Mix file (.gm)")
	yellowColor.Println("  go-mix server <port>      Start REPL server on specified port")
	yellowColor.Println("  go-mix --help             Display this help message")
	yellowColor.Println("  go-mix --version          Display version information")
	cyanColor.Println("")
	cyanColor.Println("REPL COMMANDS:")
	yellowColor.Println("  /exit                     Exit the REPL")
	yellowColor.Println("  /scope                    Show current scope and variables")
	cyanColor.Println("")
	cyanColor.Println("EXAMPLES:")
	yellowColor.Println("  go-mix                    # Start REPL")
	yellowColor.Println("  go-mix samples/algo/05_factorial.gm")
	yellowColor.Println("  go-mix server 8080        # Start REPL server on port 8080")
	cyanColor.Println("")
	cyanColor.Println("For more information, visit: https://github.com/akashmaji946/go-mix")
}

// showVersion displays the version information for the Go-Mix interpreter
func showVersion() {
	cyanColor.Println("Go-Mix - An Interpreted Programming Language")
	cyanColor.Printf("Version: %s\n", VERSION)
	cyanColor.Printf("License: %s\n", LICENCE)
	cyanColor.Printf("Author : %s\n", AUTHOR)
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

// startServer initializes and runs the Go-Mix REPL server.
// It listens on the specified port for incoming TCP connections.
// Each connection is handled in a separate goroutine, providing a dedicated REPL session.
//
// Parameters:
//
//	port - The network port to listen on (e.g., "8080")
func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		redColor.Fprintf(os.Stderr, "[SERVER ERROR] Failed to start server on port %s: %v\n", port, err)
		os.Exit(1)
	}
	cyanColor.Printf("Go-Mix REPL server listening on :%s\n", port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			redColor.Fprintf(os.Stderr, "[SERVER ERROR] Failed to accept connection: %v\n", err)
			continue
		}
		go handleClient(conn)
	}
}

// handleClient manages a single client connection for the REPL server.
// It creates a new REPL instance and starts it, using the network connection
// as both the input reader and output writer.
func handleClient(conn net.Conn) {
	defer conn.Close()
	cyanColor.Printf("New client connected from %s\n", conn.RemoteAddr())
	repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE, PROMPT)
	repler.Start(conn, conn) // Use the network connection as stdin/stdout
	cyanColor.Printf("Client disconnected from %s\n", conn.RemoteAddr())
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

	// Print the AST for debugging purposes (currently commented out for cleaner output)
	// fmt.Println("Parsed AST:")
	// printAST(rootNode)

	// Create evaluator and execute the AST
	// The evaluator walks the AST and executes the program
	evaluator := eval.NewEvaluator()
	evaluator.SetParser(par) // Link parser for access to environment and error handling
	result := evaluator.Eval(rootNode)

	// Display result if any (and not nil)
	// Only display non-nil results to avoid cluttering output
	if result != nil {
		if result.GetType() == "error" {
			// Evaluation produced an error object - display and exit
			redColor.Fprintf(os.Stderr, "%s\n", result.ToString())
			os.Exit(1)
		} else {
			// Successful evaluation - display result in yellow
			if result.GetType() != "nil" { // Skip printing null results for cleaner output
				yellowColor.Fprintf(os.Stdout, "%s\n", result.ToString())
			}
		}
	}
}

// printAST is a helper function to display the AST structure for debugging.
// It recursively prints the AST nodes with indentation to show hierarchy.
//
// Parameters:
//   - node: The root AST node to print
//   - indent: The current indentation level (used for recursive calls)
//
// This function is useful for developers to visualize the parsed structure of the code.
// It can be enabled during development and debugging to understand how the parser is interpreting the source code.
func printAST(rootNode *parser.RootNode) {
	p := &PrintingVisitor{}
	p.VisitRootNode(*rootNode)
	fmt.Println(p.Buf.String())
}
