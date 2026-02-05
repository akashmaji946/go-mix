/*
File    : go-mix/repl/repl.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)

Package repl implements the Read-Eval-Print Loop (REPL) for the Go-Mix interpreter.
The REPL provides an interactive environment where users can:
- Enter Go-Mix code line by line
- See immediate results of their code execution
- Navigate command history using arrow keys
- Receive colored feedback for different types of output

The REPL uses the readline library for enhanced line editing capabilities
and integrates with the parser and evaluator to execute user input.
*/
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
// These colors provide visual feedback to enhance user experience:
// - blueColor: Decorative lines and separators
// - yellowColor: Expression results and version info
// - redColor: Error messages and warnings
// - greenColor: Banner and success messages
// - cyanColor: Informational messages and instructions
var (
	blueColor   = color.New(color.FgBlue)
	yellowColor = color.New(color.FgYellow)
	redColor    = color.New(color.FgRed)
	greenColor  = color.New(color.FgGreen)
	cyanColor   = color.New(color.FgCyan)
)

// Repl represents the Read-Eval-Print Loop instance.
// It encapsulates all the configuration needed to run an interactive session.
type Repl struct {
	Banner  string // ASCII art banner displayed at startup
	Version string // Version string of the interpreter
	Author  string // Author contact information
	Line    string // Separator line for visual formatting
	License string // Software license information
	Prompt  string // Command prompt shown to the user (e.g., "gm >>> ")
}

// NewRepl creates and initializes a new REPL instance.
// This constructor sets up all the visual elements and configuration
// needed for the interactive session.
//
// Parameters:
//
//	banner  - ASCII art logo to display at startup
//	version - Version string of the interpreter
//	author  - Author contact information
//	line    - Separator line for formatting
//	license - Software license information
//	prompt  - Command prompt string
//
// Returns:
//
//	A pointer to a newly created Repl instance
func NewRepl(banner string, version string, author string, line string, license string, prompt string) *Repl {
	return &Repl{Banner: banner, Version: version, Author: author, Line: line, License: license, Prompt: prompt}
}

// PrintBannerInfo displays the welcome banner and usage instructions.
// This function is called when the REPL starts to provide users with:
// - The Go-Mix logo (ASCII art)
// - Version and author information
// - Basic usage instructions
// - Command history navigation tips
//
// The output uses colors to make the information visually appealing
// and easy to read.
//
// Parameters:
//
//	writer - The io.Writer to output the banner to (typically os.Stdout)
func (r *Repl) PrintBannerInfo(writer io.Writer) {

	// Print top separator line in blue
	blueColor.Fprintf(writer, "%s\n", r.Line)

	// Print the ASCII art banner in green
	greenColor.Fprintf(writer, "%s\n", r.Banner)

	// Print separator line
	blueColor.Fprintf(writer, "%s\n", r.Line)

	// Print version, author, and license information in yellow
	yellowColor.Fprintln(writer, "Version: "+r.Version+" | Author: "+r.Author+" | Lincense: "+r.License)

	// Print separator line
	blueColor.Fprintf(writer, "%s\n", r.Line)

	// Print welcome message and usage instructions in cyan
	cyanColor.Fprintf(writer, "%s\n", "Welcome to Go-Mix!")
	cyanColor.Fprintf(writer, "%s\n", "Type your code and press enter")
	cyanColor.Fprintf(writer, "%s\n", "Type '.exit' to quit")
	cyanColor.Fprintf(writer, "%s\n", "Use up/down arrows to navigate command history")

	// Print bottom separator line
	blueColor.Fprintf(writer, "%s\n", r.Line)
}

// Start begins the REPL main loop.
// This is the core function that handles the interactive session:
// 1. Displays the welcome banner
// 2. Sets up readline for line editing and history
// 3. Creates an evaluator instance
// 4. Enters the main read-eval-print loop
// 5. Processes user input until exit
//
// The loop continues until:
// - User types '.exit'
// - EOF is encountered (Ctrl+D)
// - An error occurs in readline
//
// Parameters:
//
//	reader - Input source (typically os.Stdin, though not directly used due to readline)
//	writer - Output destination (typically os.Stdout)
//
// Features:
// - Command history (accessible via up/down arrows)
// - Line editing capabilities (backspace, cursor movement, etc.)
// - Automatic whitespace trimming
// - Empty line handling
// - Panic recovery for robust error handling
func (r *Repl) Start(reader io.Reader, writer io.Writer) {

	// Print the welcome banner and usage instructions
	r.PrintBannerInfo(writer)

	// Create a new readline instance for enhanced line editing
	// This provides features like command history, cursor movement, etc.
	rl, err := readline.New(r.Prompt)
	if err != nil {
		panic(err)
	}
	defer rl.Close() // Ensure readline is properly closed on exit

	// Create a new evaluator instance for executing Go-Mix code
	evaluator := eval.NewEvaluator()
	evaluator.SetWriter(writer) // Set output writer for print statements

	// Main REPL loop - continues until user exits or error occurs
	for {
		// Read a line of input from the user
		// This blocks until the user presses Enter
		line, err := rl.Readline()
		if err != nil {
			// EOF or error occurred (e.g., Ctrl+D pressed)
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		// Trim whitespace from the input
		line = strings.Trim(line, " \n\t\r")

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check for exit command
		if line == ".exit" {
			writer.Write([]byte("Good Bye!\n"))
			break
		}

		// Save the command to history for up/down arrow navigation
		rl.SaveHistory(line)

		// Echo input in blue (currently commented out)
		// blueColor.Fprintf(writer, ">> %s\n", line)

		// Execute the input with panic recovery to prevent crashes
		r.executeWithRecovery(writer, line, evaluator)
	}
}

// executeWithRecovery handles parsing and evaluation with panic recovery.
// This function implements a robust error handling strategy for the REPL:
// 1. Sets up panic recovery to catch runtime errors
// 2. Parses the user input into an AST
// 3. Checks for parsing errors
// 4. Evaluates the AST
// 5. Displays results or errors
//
// Unlike file execution mode, the REPL continues running after errors,
// allowing users to correct mistakes and try again.
//
// Parameters:
//
//	writer    - Output destination for results and errors
//	line      - The user's input line to execute
//	evaluator - The evaluator instance (maintains state across REPL sessions)
//
// Error Handling:
//   - Panics: Caught and displayed as runtime errors, REPL continues
//   - Parse errors: Displayed in red, REPL continues
//   - Evaluation errors: Displayed in red, REPL continues
//   - Success: Result displayed in yellow
func (r *Repl) executeWithRecovery(writer io.Writer, line string, evaluator *eval.Evaluator) {
	// Recover from any panics that might occur during parsing or evaluation
	// Unlike file mode, we don't exit - just display the error and continue
	defer func() {
		if recovered := recover(); recovered != nil {
			redColor.Fprintf(writer, "[RUNTIME ERROR] %v\n", recovered)
		}
	}()

	// Parse the input line into an Abstract Syntax Tree (AST)
	par := parser.NewParser(line)
	rootNode := par.Parse()

	// Check for parser errors
	// The parser collects errors instead of panicking
	if par.HasErrors() {
		for _, err := range par.GetErrors() {
			redColor.Fprintf(writer, "%s\n", err)
		}
		return // Return to REPL prompt for user to try again
	}

	// Verify that parsing produced a valid AST root node
	if rootNode == nil {
		redColor.Fprintf(writer, "[LEXER ERROR] Invalid syntax or parser error\n")
		return // Return to REPL prompt
	}

	// Link the parser to the evaluator for access to environment
	evaluator.SetParser(par)

	// Evaluate the AST and get the result
	result := evaluator.Eval(rootNode)

	// Display the result if it's not nil
	if result != nil {
		if result.GetType() == "error" {
			// Evaluation produced an error - display in red
			redColor.Fprintf(writer, "%s\n", result.ToString())
		} else {
			// Successful evaluation - display result in yellow
			// Note: nil results are still printed (unlike file mode)
			yellowColor.Fprintf(writer, "%s\n", result.ToString())
		}
	}
}
